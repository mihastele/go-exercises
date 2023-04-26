package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func wait() {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
}

type Hits struct {
	hits       int
	sync.Mutex // embedded mutex to the structure
}

func serve(wg *sync.WaitGroup, hits *Hits, iteration int) {
	wait()
	hits.Lock()
	defer hits.Unlock()
	defer wg.Done()
	hits.hits += 1
	println("served iteration:", iteration)

}

func main() {
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup
	hitCounter := Hits{} // default values
	for i := 0; i < 30; i++ {
		iteration := i
		wg.Add(1)
		go serve(&wg, &hitCounter, iteration)
	}

	fmt.Printf("Waiting for %d goroutines to finish\n", 30)
	wg.Wait()

	hitCounter.Lock()
	totalHits := hitCounter.hits
	hitCounter.Unlock()
	fmt.Println("Total hits:", totalHits)
}
