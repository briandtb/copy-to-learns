package api

import (
	"time"

	"k8s.io/metrics/pkg/apis/metrics"
	corev1 "k8s.io/api/core/v1"
	apitypes "k8s.io/apimachinery/pkg/types"
)

// MetricsGetter is both a PodMetricGetter and a NodeMetricsGetter
type MetricsGetter interface {
	PodMetricsGetter
	NodeMetricsGetter
}

// TimeInfo represents the timing information for a metric, which was
// potentially calculated over some window of time (e.g. for CPU usage rate).
type TimeInfo struct {
	// NB: we consider the earliest timestamp amongst multiple containers
	// for the purposes of determining if a metric is tained by a time
	// period, like pod startup (used by things like the HPA).

	// Timestamp is the time at which the metrics were initially collected.
	// In the case of a rate metric, it should be the timestamp of the last
	// data point used in the calculation. If it represents multiple metric
	// points, it should be the earliest such timestamp from all of the points.
	Timestamp time.Time

	// Window represents the window used to calculate rate metrics associated 
	// with this timestamp.
	Window time.Duration
}

// PodMetricsGetter knows how to fetch metrics for the containers in a pod.
type PodMetricsGetter interface {
	// GetContainerMetrics gets the latest metrics for all containers in each listed pod,
	// returning both the metrics and the associated collection timestamp.
	// If a pod is missing, the container metrics should be nil for that pod.
	GetContainerMetrics(pods ...apitypes.NamespacedName) ([]TimeInfo, [][]metrics.ContainerMetrics)
}

// NodeMetricsGetter knows how to fetch metrics for a node.
type NodeMetricsGetter interface {
	// GetNodeMetrics gets the latest metrics for the given nodes,
	// returning both the metrics and the associated collection timestamp.
	// If a node is missing, the resourcelist should be nil for that node.
	GetNodeMetrics(nodes ...string) ([]TimeInfo, []corev1.ResourceList)
}
