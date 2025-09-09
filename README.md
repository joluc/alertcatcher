# AlertCatcher

A simple webhook receiver for Prometheus Alertmanager alerts.

## Features

- Receives and processes Alertmanager webhook notifications
- Exposes Prometheus metrics for monitoring
- Health check endpoint
- Kubernetes deployment ready

## Endpoints

- `POST /api/v1/webhook` - Webhook endpoint for Alertmanager
- `GET /health` - Health check endpoint
- `GET /metrics` - Prometheus metrics

## Running Locally

```bash
# Build and run
go build -o alertcatcher .
./alertcatcher

# Or run directly
go run main.go
```

The server will start on port 8080 by default. You can override this with the `PORT` environment variable.

## Docker

```bash
# Build image
docker build -t alertcatcher .

# Run container
docker run -p 8080:8080 alertcatcher
```

## Kubernetes Deployment

Deploy using the provided Helm chart:

```bash
helm install alertcatcher ./chart/alertcatcher
```

## Configuration

The application uses environment variables for configuration:

- `PORT` - Server port (default: 8080)

## Metrics

The following Prometheus metrics are exposed:

- `alertcatcher_incoming_requests_total` - Total number of incoming webhook requests
- `alertcatcher_processed_alerts_total` - Total number of processed alerts

## Testing

Send a test webhook:

```bash
curl -X POST http://localhost:8080/api/v1/webhook \
  -H "Content-Type: application/json" \
  -d '{"alerts":[{"status":"firing","labels":{"alertname":"test"}}]}'
```
