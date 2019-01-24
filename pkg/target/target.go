package target

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Target interface {
	Collect(chan<- prometheus.Metric)
	Describe(chan<- *prometheus.Desc)
	Run() error
}
