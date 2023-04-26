package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup
	LengthOfConcurrent := 5
	counter := LengthOfConcurrent
	fmt.Println("Waiting for goroutines to finish")
	for i := 0; i < LengthOfConcurrent; i++ {
		wg.Add(1)

		go func() {
			defer func() {
				defer func() {
					fmt.Println("goroutines remaining =", counter)
					counter -= 1
					wg.Done()
				}()
				duration := time.Duration(rand.Intn(1000)) * time.Millisecond
				fmt.Println("Waiting for", duration)
				time.Sleep(duration)
			}()
		}()
	}

	wg.Wait()
	fmt.Println("All goroutines finished")
}
