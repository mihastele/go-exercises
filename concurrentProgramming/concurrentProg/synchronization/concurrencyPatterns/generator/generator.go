package main

import (
	"fmt"
	"math/rand"
	"time"
)

// infinite generator
func generateRandInt(min, max int) <-chan int {
	out := make(chan int, 3)

	go func() {
		for {
			out <- rand.Intn(max-min+1) + min
		}
	}()

	return out
}

// finite generator
func generateRandInt2(count, min, max int) <-chan int {
	out := make(chan int, 1)

	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Intn(max-min+1) + min
		}
		close(out)
	}()

	return out
}

func main() {
	rand.Seed(time.Now().UnixNano())

	randInt := generateRandInt(0, 100)

	fmt.Println("Generated random int: ")
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)

	randInt1 := generateRandInt2(2, 0, 100)

	fmt.Println("Generated random int: ")
	fmt.Println(<-randInt1)
	fmt.Println(<-randInt1)
	fmt.Println(<-randInt1)
	fmt.Println(<-randInt1)

	// iterate through generator
	randIntRange := generateRandInt2(3, 1, 10)
	for {
		n, open := <-randIntRange
		if !open {
			break
		}
		fmt.Println(n)
	}
}
