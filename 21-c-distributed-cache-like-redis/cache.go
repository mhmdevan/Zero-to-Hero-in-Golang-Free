package main

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

type Cache struct {
	data map[string]string
	mu   sync.RWMutex
	file string // File for persistent storage
}

func NewCache(file string) *Cache {
	c := &Cache{
		data: make(map[string]string),
		file: file,
	}
	c.loadFromDisk()
	return c
}

func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
	c.saveToDisk()
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, found := c.data[key]
	return val, found
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
	c.saveToDisk()
}

func (c *Cache) saveToDisk() {
	file, err := os.Create(c.file)
	if err != nil {
		log.Println("Error saving to disk:", err)
		return
	}
	defer file.Close()
	json.NewEncoder(file).Encode(c.data)
}

func (c *Cache) loadFromDisk() {
	file, err := os.Open(c.file)
	if err != nil {
		log.Println("No existing data found, starting fresh.")
		return
	}
	defer file.Close()
	json.NewDecoder(file).Decode(&c.data)
}
