package storage

import "k8s.io/component-base/metrics"

var (
	pointsStored = metrics.NewGaugeVec(
		&metrics.GaugeOpts{
			Namespace: "metrics_server",
			Subsystem: "storage",
			Name:      "points",
			Help:      "Number of metrics points stored.",
		},
		[]string{"type"},
	)
)

// RegisterStorageMetrics registers a gauge metric for the number of metrics
// points stored.
func RegisterStorageMetrics(registrationFunc func(metrics.Registerable) error) error {
	return registrationFunc(pointsStored)
}
