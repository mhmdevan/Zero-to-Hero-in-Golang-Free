package main

import (
	"log"
	"net/http"
	"time"
)

func (hr *HashRing) MonitorNodes() {
	for node := range hr.nodeMapping {
		go func(node string) {
			for {
				resp, err := http.Get("http://" + node + "/health")
				if err != nil || resp.StatusCode != http.StatusOK {
					log.Printf("Node %s is down. Removing from hash ring.", node)
					hr.RemoveNode(node)
					break
				}
				time.Sleep(5 * time.Second)
			}
		}(node)
	}
}
