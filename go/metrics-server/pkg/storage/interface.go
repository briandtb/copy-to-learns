package storage

import "metrics-server/pkg/api"

type Storage interface {
	api.MetricsGetter
	Store(batch *MetricsBatch)
}
