package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	cacheHits      = prometheus.NewCounter(prometheus.CounterOpts{Name: "cache_hits_total", Help: "Total cache hits"})
	cacheMisses    = prometheus.NewCounter(prometheus.CounterOpts{Name: "cache_misses_total", Help: "Total cache misses"})
	requestLatency = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "cache_request_latency_seconds",
		Help:    "Latency of cache requests",
		Buckets: prometheus.DefBuckets,
	}, []string{"operation"})
	requestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cache_requests_total",
		Help: "Total requests received",
	}, []string{"operation"})
)

func init() {
	prometheus.MustRegister(cacheHits, cacheMisses, requestLatency, requestsTotal)
}

func recordCacheHit() {
	cacheHits.Inc()
	requestsTotal.WithLabelValues("hit").Inc()
}

func recordCacheMiss() {
	cacheMisses.Inc()
	requestsTotal.WithLabelValues("miss").Inc()
}

func recordLatency(operation string, duration float64) {
	requestLatency.WithLabelValues(operation).Observe(duration)
}
