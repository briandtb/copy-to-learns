package storage

import (
	"time"

	"k8s.io/apimachinery/pkg/api/resource"
)

// MetricsBatch is a single batch of pod, container, and node metrics from some source.
type MetricsBatch struct {
	Nodes []NodeMetricsPoint
	Pods  []PodMetricsPoint
}

// NodeMetricsPoint contains the metrics for some node at some point in time.
type NodeMetricsPoint struct {
	Name string
	MetricsPoint
}

// PodMetricsPoint contains the metrics for some pod's containers.
type PodMetricsPoint struct {
	Name      string
	Namespace string

	Containers []ContainerMetricsPoint
}

// ContainerMetricsPoint contains the metrics for some container at some point in time.
type ContainerMetricsPoint struct {
	Name string
	MetricsPoint
}

// MetricsPoint represents a set of specific metrics at some point in time.
type MetricsPoint struct {
	Timestamp time.Time
	// CpuUsage is the CPU usage rate, in cores.
	CpuUsage resource.Quantity
	// MemoryUsage is the working set size, in bytes.
	MemoryUsage resource.Quantity
}
