package main

import (
	"fmt"
	"time"
)

// asynchrounous and threaded
// go routines

func count(amount int) {
	for i := 0; i < amount; i++ {
		time.Sleep(100 * time.Millisecond)
		println(i)
	}
}

func main() {
	go count(10)
	fmt.Println("wait for goroutine")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("done")
}
