package main

import (
	"fmt"
	"time"
	"unicode"
)

func main() {
	counter := 0
	wait := func(ms time.Duration) {
		time.Sleep(ms * time.Millisecond)
		counter += 1
	}

	fmt.Println("wait for goroutine")
	go wait(100)
	go wait(900)
	go wait(1000)

	fmt.Println("Launched. Counter =", counter)
	time.Sleep(1500 * time.Millisecond)
	fmt.Println("Done. Counter =", counter)

	data := []rune{'a', 'b', 'c', 'd'}
	var capitalized []rune
	var capts [][]rune

	capIt := func(r rune, capitalized *[]rune) {
		*capitalized = append(*capitalized, unicode.ToUpper(r))
		fmt.Printf("%c done!\n", r)
	}

	for ccc := 0; ccc < 40; ccc++ {
		capitalized = []rune{}
		for i := 0; i < len(data); i++ {
			go capIt(data[i], &capitalized)
		}
		capts = append(capts, capitalized)
	}

	time.Sleep(10000 * time.Millisecond)
	fmt.Printf("Capitalized: %c\n", capts)
	// you can tell that from goroutines that output is not deterministic due to concurrency
}
