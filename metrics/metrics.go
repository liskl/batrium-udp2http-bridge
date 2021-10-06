package metrics

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-kit/kit/metrics"
)

// NewMetrics construct metrics builder with included metrics provider.
func NewMetrics(driver Driver) *MetricController {
	return &MetricController{
		driver:     driver,
		gauges:     make(map[string]metrics.Gauge),
		counters:   make(map[string]metrics.Counter),
		histograms: make(map[string]metrics.Histogram),
	}
}

// MetricController – metrics controller.
type MetricController struct {
	driver     Driver
	counters   map[string]metrics.Counter
	histograms map[string]metrics.Histogram
	gauges     map[string]metrics.Gauge
}

func getMetricName(m string, t []string) string {
	return fmt.Sprintf("%s|%s", m, strings.Join(t, ","))
}

// Counter – update counter value.
func (mc *MetricController) Counter(desc, metric string, tags []string, count float64) {
	storedMetric := getMetricName(metric, tagsIntoTagNames(tags))
	if mc.counters[storedMetric] == nil {
		counter := mc.driver.NewCounter(desc, metric, tags)
		mc.counters[storedMetric] = counter
	}
	mc.counters[storedMetric].With(tags...).Add(count)
}

// Elapsed reports the time elapsed for a function call in seconds
func (mc *MetricController) Elapsed(desc string, metric string, tags []string, buckets ...float64) func(time.Time) {
	return func(begin time.Time) {
		mc.Histogram(desc, metric, tags, time.Since(begin).Seconds(), buckets...)
	}
}

// Histogram measure the statistical distribution of a set of values
func (mc *MetricController) Histogram(desc, metric string, tags []string, value float64, buckets ...float64) {
	storedMetric := getMetricName(metric, tagsIntoTagNames(tags))
	if mc.histograms[storedMetric] == nil {
		histogram := mc.driver.NewHistogram(desc, metric, tags, buckets...)
		mc.histograms[storedMetric] = histogram
	}
	mc.histograms[storedMetric].With(tags...).Observe(value)
}

func (mc *MetricController) getOrCreateGauge(desc, metric string, tags []string) metrics.Gauge {
	storedMetric := getMetricName(metric, tagsIntoTagNames(tags))
	if mc.gauges[storedMetric] == nil {
		gauge := mc.driver.NewGauge(desc, metric, tags)
		mc.gauges[storedMetric] = gauge
	}
	return mc.gauges[storedMetric].With(tags...)
}

// GaugeSet set gauge value.
func (mc *MetricController) GaugeSet(desc, metric string, tags []string, value float64) {
	mc.getOrCreateGauge(desc, metric, tags).Set(value)
}

// GaugeAdd update gauge value.
func (mc *MetricController) GaugeAdd(desc, metric string, tags []string, value float64) {
	mc.getOrCreateGauge(desc, metric, tags).Add(value)
}

func tagsIntoTagNames(tags []string) []string {
	var tagNames []string
	for i := 0; i < len(tags); i++ {
		if i%2 == 0 {
			tagNames = append(tagNames, tags[i])
		}
	}
	return tagNames
}
