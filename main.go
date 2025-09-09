package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joluc/alertcatcher/pkg/handler"
	"github.com/joluc/alertcatcher/pkg/metrics"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.Handle("/metrics", metrics.MetricsHandler())
	http.HandleFunc("/health", handler.HealthCheck)
	http.HandleFunc("/api/v1/webhook", func(w http.ResponseWriter, r *http.Request) {
		metrics.IncomingRequests.Inc()
		handler.HandleWebhook(w, r)
	})

	log.Printf("AlertCatcher starting on port %s...", port)
	log.Printf("Health check available at: /health")
	log.Printf("Metrics available at: /metrics")
	log.Printf("Webhook endpoint: /api/v1/webhook")

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
