package db

import (
	"github.com/go-kit/kit/metrics/discard"
	"github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

// metrics
var (
	metricUncloseItr = discard.NewGauge()
)

func init() {

	metricUncloseItr = prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
		Namespace: "db",
		Subsystem: "stats",
		Name:      "iterators_unclosed",
		Help:      "number of unclosed iterators",
	}, []string{})

}
