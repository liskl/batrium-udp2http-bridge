package metrics

import (
	"time"

	"github.com/go-kit/kit/metrics"
)

// Metrics – To collect metric values.
type Metrics interface {
	Elapsed(desc string, metric string, tags []string, buckets ...float64) func(time.Time)
	Counter(desc string, metric string, tags []string, value float64)
	Histogram(desc string, metric string, tags []string, value float64, buckets ...float64)
	GaugeSet(desc string, metric string, tags []string, value float64)
	GaugeAdd(desc string, metric string, tags []string, value float64)
}

// Driver – Construct metric collectors.
type Driver interface {
	NewCounter(string, string, []string) metrics.Counter
	NewHistogram(string, string, []string, ...float64) metrics.Histogram
	NewGauge(string, string, []string) metrics.Gauge
}