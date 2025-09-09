package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/joluc/alertcatcher/pkg/metrics"
)

// HandleWebhook processes incoming webhook requests from Alertmanager.
func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is accepted", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	log.Printf("Received webhook data from %s", r.RemoteAddr)

	var alert AlertData
	if err := json.Unmarshal(body, &alert); err != nil {
		log.Printf("Error unmarshaling JSON: %v", err)
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Process the alert data
	processAlert(alert)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"success","message":"Webhook processed successfully"}`))

	log.Printf("Successfully processed %d alerts", len(alert.Alerts))
}

func processAlert(alert AlertData) {
	log.Printf("Processing %d alerts from receiver: %s", len(alert.Alerts), alert.Receiver)

	for i, a := range alert.Alerts {
		log.Printf("Alert %d: Status=%s, Summary=%s", i+1, a.Status, a.Annotations["summary"])
		metrics.ProcessedAlerts.Inc()
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"healthy","service":"alertcatcher"}`))
}
