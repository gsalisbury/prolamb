package prolamb

import (
	"net/http"
	"net/url"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/tlmiller/prolamb/pkg/config"
)

const (
	TargetParam = "target"
)

func PrometheusHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	vals, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		panic(err)
	}

	target, err := config.GetTargetsFromRawQuery(vals.Get(TargetParam))
	if err := target.Run(); err != nil {
		panic(err)
	}
	register := prometheus.NewRegistry()
	register.MustRegister(target)
	handle := promhttp.HandlerFor(register, promhttp.HandlerOpts{})
	handle.ServeHTTP(w, r)
}
