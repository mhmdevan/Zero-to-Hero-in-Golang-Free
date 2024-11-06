package main

import (
	"hash/crc32"
	"sort"
	"strconv"
)

type HashRing struct {
	nodes             map[int]string // Virtual nodes
	nodeKeys          []int          // Sorted keys for the ring
	replicas          int            // Number of virtual nodes per node
	replicationFactor int            // Number of replicas for each key
	nodeMapping       map[string]*Cache
}

func NewHashRing(replicas, replicationFactor int) *HashRing {
	return &HashRing{
		nodes:             make(map[int]string),
		nodeKeys:          []int{},
		replicas:          replicas,
		replicationFactor: replicationFactor,
		nodeMapping:       make(map[string]*Cache),
	}
}

func (hr *HashRing) AddNode(node string, cache *Cache) {
	for i := 0; i < hr.replicas; i++ {
		vnodeKey := int(crc32.ChecksumIEEE([]byte(node + strconv.Itoa(i))))
		hr.nodes[vnodeKey] = node
		hr.nodeKeys = append(hr.nodeKeys, vnodeKey)
	}
	sort.Ints(hr.nodeKeys)
	hr.nodeMapping[node] = cache
}

func (hr *HashRing) RemoveNode(node string) {
	for i := 0; i < hr.replicas; i++ {
		vnodeKey := int(crc32.ChecksumIEEE([]byte(node + strconv.Itoa(i))))
		delete(hr.nodes, vnodeKey)
	}
	// Update nodeKeys and remove from map
	newKeys := []int{}
	for _, key := range hr.nodeKeys {
		if hr.nodes[key] != node {
			newKeys = append(newKeys, key)
		}
	}
	hr.nodeKeys = newKeys
	delete(hr.nodeMapping, node)
}

func (hr *HashRing) GetCaches(key string) []*Cache {
	if len(hr.nodeKeys) == 0 {
		return nil
	}
	vnodeKey := int(crc32.ChecksumIEEE([]byte(key)))
	idx := sort.Search(len(hr.nodeKeys), func(i int) bool {
		return hr.nodeKeys[i] >= vnodeKey
	})
	var caches []*Cache
	for i := 0; i < hr.replicationFactor; i++ {
		if idx >= len(hr.nodeKeys) {
			idx = 0
		}
		node := hr.nodes[hr.nodeKeys[idx]]
		caches = append(caches, hr.nodeMapping[node])
		idx++
	}
	return caches
}
