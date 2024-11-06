package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type CacheService struct {
	ring *HashRing
}

func NewCacheService(replicas, replicationFactor int, nodes []string) *CacheService {
	ring := NewHashRing(replicas, replicationFactor)
	for i, node := range nodes {
		cache := NewCache("node_" + strconv.Itoa(i) + ".json") // Persistent storage per node
		ring.AddNode(node, cache)
	}
	return &CacheService{ring: ring}
}

func (cs *CacheService) setHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	caches := cs.ring.GetCaches(req.Key)
	if len(caches) == 0 {
		http.Error(w, "No cache available", http.StatusInternalServerError)
		return
	}

	// Quorum write: requires majority of replicas to acknowledge
	ackCount := 0
	for _, cache := range caches {
		cache.Set(req.Key, req.Value)
		ackCount++
		if ackCount >= len(caches)/2+1 { // Quorum threshold
			break
		}
	}

	if ackCount < len(caches)/2+1 {
		http.Error(w, "Failed to achieve write quorum", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	recordCacheMiss() // Prometheus metric
}

func (cs *CacheService) getHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	caches := cs.ring.GetCaches(key)
	if len(caches) == 0 {
		http.Error(w, "No cache available", http.StatusInternalServerError)
		return
	}

	readCount := 0
	var value string
	found := false
	for _, cache := range caches {
		if v, ok := cache.Get(key); ok {
			value = v
			readCount++
			if readCount >= len(caches)/2+1 { // Quorum threshold
				found = true
				break
			}
		}
	}

	if !found {
		http.Error(w, "Key not found", http.StatusNotFound)
		recordCacheMiss()
		return
	}

	recordCacheHit()
	json.NewEncoder(w).Encode(map[string]string{"value": value})
}

func (cs *CacheService) deleteHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	caches := cs.ring.GetCaches(key)
	if len(caches) == 0 {
		http.Error(w, "No cache available", http.StatusInternalServerError)
		return
	}
	for _, cache := range caches {
		cache.Delete(key)
	}
	w.WriteHeader(http.StatusOK)
}

func (cs *CacheService) joinHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Node string `json:"node"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	newCache := NewCache(req.Node + ".json")
	cs.ring.AddNode(req.Node, newCache)
	w.WriteHeader(http.StatusOK)
}

func (cs *CacheService) leaveHandler(w http.ResponseWriter, r *http.Request) {
	node := r.URL.Query().Get("node")
	cs.ring.RemoveNode(node)
	w.WriteHeader(http.StatusOK)
}

func (cs *CacheService) healthHandler(w http.ResponseWriter, r *http.Request) {
	status := map[string]interface{}{
		"status": "OK",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}

func main() {
	nodes := []string{"node1", "node2", "node3"}
	service := NewCacheService(3, 2, nodes) // 3 replicas, replication factor of 2

	http.HandleFunc("/set", service.setHandler)
	http.HandleFunc("/get", service.getHandler)
	http.HandleFunc("/delete", service.deleteHandler)
	http.HandleFunc("/join", service.joinHandler)
	http.HandleFunc("/leave", service.leaveHandler)
	http.HandleFunc("/health", service.healthHandler)
	http.Handle("/metrics", promhttp.Handler()) // Prometheus metrics endpoint

	service.ring.MonitorNodes() // Start monitoring nodes for health

	log.Println("Distributed Cache Service is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
