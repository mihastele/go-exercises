package main

import (
	"fmt"
	"sync"
	"time"
)

type SyncedData struct {
	inner map[string]int
	mutex sync.Mutex
}

func (sd *SyncedData) Insert(key string, value int) {
	sd.mutex.Lock()
	sd.inner[key] = value
	sd.mutex.Unlock()
}

func (sd *SyncedData) Get(key string) int {
	sd.mutex.Lock()
	data := sd.inner[key]
	sd.mutex.Unlock()
	return data
}

func main() {
	data := SyncedData{inner: make(map[string]int)}
	data.Insert("a", 1)
	data.Insert("b", 2)
	go fmt.Println(data.Get("a"))
	go fmt.Println(data.Get("b"))
	time.Sleep(100 * time.Millisecond)
}
