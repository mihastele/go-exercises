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

// defer  can make sure the mutex gets unlocked
func (sd *SyncedData) InsertDefer(key string, value int) {
	sd.mutex.Lock()
	defer sd.mutex.Unlock()
	sd.inner[key] = value
}

func (sd *SyncedData) GetDefer(key string) int {
	sd.mutex.Lock()
	defer sd.mutex.Unlock()
	data := sd.inner[key]
	return data
}

//func main() {
//	data := SyncedData{inner: make(map[string]int)}
//	data.Insert("a", 1)
//	data.Insert("b", 2)
//	go fmt.Println(data.Get("a"))
//	go fmt.Println(data.Get("b"))
//	time.Sleep(100 * time.Millisecond)
//}

func main() {
	data := SyncedData{inner: make(map[string]int)}
	data.Insert("a", 1)
	data.Insert("b", 2)
	fmt.Println(data.Get("a"))
	fmt.Println(data.Get("b"))
	time.Sleep(100 * time.Millisecond)
}
