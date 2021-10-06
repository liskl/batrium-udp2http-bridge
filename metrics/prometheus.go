package metrics

import (
	"github.com/go-kit/kit/metrics"
	prometheus2 "github.com/go-kit/kit/metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus"
)

// NewPrometheusMetrics – provide prometheus metrics
func NewPrometheusMetrics() Metrics {
	driver := &pd{}
	m := NewMetrics(driver)
	StartPrometheusServer("")
	return m
}

type pd struct {
}

// NewCounter – constructor for counter.
func (d *pd) NewCounter(desc string, metric string, tags []string) metrics.Counter {
	vec := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: metric,
			Help: desc,
		},
		getLabels(tags))
	prometheus.MustRegister(vec)
	return prometheus2.NewCounter(vec)
}

// NewHistogram – constructor for histogram.
func (d *pd) NewHistogram(desc string, metric string, tags []string, buckets ...float64) metrics.Histogram {
	vec := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    metric,
			Help:    desc,
			Buckets: buckets,
		},
		getLabels(tags))
	prometheus.MustRegister(vec)
	return prometheus2.NewHistogram(vec)
}

// NewGauge – constructor for gauge.
func (d *pd) NewGauge(desc string, metric string, tags []string) metrics.Gauge {
	vec := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: metric,
			Help: desc,
		},
		getLabels(tags))
	prometheus.MustRegister(vec)
	return prometheus2.NewGauge(vec)
}

func getLabels(tags []string) (res []string) {
	for i := 0; i < len(tags); i = i + 2 {
		res = append(res, tags[i])
	}
	return
}
