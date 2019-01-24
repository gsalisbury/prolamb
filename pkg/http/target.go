package http

import (
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"

	"github.com/prometheus/client_golang/prometheus"
)

type Target struct {
	URL     *url.URL
	Method  string
	Name    string
	metrics []prometheus.Metric
}

var (
	HTTPRequestTime *prometheus.Desc
)

func (t *Target) Collect(out chan<- prometheus.Metric) {
	for _, metric := range t.metrics {
		out <- metric
	}
}

func (t *Target) Describe(out chan<- *prometheus.Desc) {
	out <- HTTPRequestTime
}

func init() {
	HTTPRequestTime = prometheus.NewDesc("http_request_time_ms", "help", []string{},
		prometheus.Labels{})
}

func (t *Target) Run() error {
	req, err := http.NewRequest(t.Method, t.URL.String(), nil)
	if err != nil {
		return errors.Wrapf(err, "making http target request for url %s", t.URL.String())
	}
	client := http.Client{}
	start := time.Now()
	_, err = client.Do(req)
	if err != nil {
		return errors.Wrapf(err, "calling http target for url %s", t.URL.String())
	}

	taken := time.Since(start)
	met, err := prometheus.NewConstMetric(HTTPRequestTime, prometheus.GaugeValue,
		float64(taken/time.Millisecond))
	if err != nil {
		return errors.Wrapf(err,
			"making http target request time metric for url %s", t.URL.String())
	}
	t.metrics = append(t.metrics, met)
	return nil
}
