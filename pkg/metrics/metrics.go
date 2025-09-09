package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	IncomingRequests = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "alertcatcher_incoming_requests_total",
			Help: "Total number of incoming webhook requests",
		},
	)

	ProcessedAlerts = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "alertcatcher_processed_alerts_total",
			Help: "Total number of processed alerts",
		},
	)
)

func init() {
	prometheus.MustRegister(IncomingRequests)
	prometheus.MustRegister(ProcessedAlerts)
}

func MetricsHandler() http.Handler {
	return promhttp.Handler()
}
