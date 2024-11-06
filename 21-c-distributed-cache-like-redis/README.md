# 🚀 Distributed Cache Service

A distributed caching service built with Go, inspired by Redis, with features including quorum-based consistency, persistent storage, dynamic scaling, and advanced monitoring with Prometheus and Grafana.

## Features
- **Quorum-Based Consistency**: Ensures data consistency across nodes using quorum reads and writes.
- **File-Based Persistent Storage**: Simple disk-based storage for data recovery after restarts.
- **Dynamic Node Management**: Nodes can dynamically join or leave the cluster.
- **Advanced Load Balancing**: Nginx is used to distribute requests evenly across multiple instances.
- **Prometheus Monitoring**: Track cache hits, misses, and node health with Prometheus metrics.
- **Grafana Dashboards**: Visualize Prometheus metrics in Grafana for insights into system performance.

---

## Project Structure

/distributed-cache │ ├── main.go # Entry point for the service, HTTP handlers ├── cache.go # Cache management with file-based persistent storage ├── consistent_hash.go # Consistent hashing for data distribution ├── node.go # Node monitoring and health checks └── metrics.go # Prometheus metrics for monitoring


---

## Setup Instructions

### Prerequisites
- **Go**: [Download Go](https://golang.org/dl/)
- **Nginx**: For load balancing.
- **Prometheus** (for monitoring): [Download Prometheus](https://prometheus.io/download/)
- **Grafana** (for dashboards): [Download Grafana](https://grafana.com/grafana/download)

### 1. Install Prometheus Client Library
```go get github.com/prometheus/client_golang/prometheus/promhttp   go mod tidy```

### 2. Run Multiple Instances of the Application
Run each instance on a different port to simulate a distributed setup:

``` go run main.go --port 8081 go run main.go --port 8082 go run main.go --port 8083 ```

### 3. Configure Nginx for Load Balancing

Add this configuration in /etc/nginx/conf.d/cache_load_balancer.conf:

``` upstream cache_nodes { server localhost:8081; server localhost:8082; server localhost:8083; } ```

``` server { listen 80; location / { proxy_pass http://cache_nodes; proxy_set_header Host $host;proxy_set_header X-Real-IP $remote_addr; proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for; proxy_set_header X-Forwarded-Proto $scheme;}} ```

Restart Nginx:
``` sudo systemctl restart nginx ```

Run Prometheus:
``` prometheus --config.file=prometheus.yml ```

Happy caching! ⚡