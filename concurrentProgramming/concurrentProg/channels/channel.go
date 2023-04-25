package main

import (
	"fmt"
	"time"
)

// same as double ended pipe (bidirectional communication)

func main() {
	channel := make(chan int)    // int communication medium used in the channel
	go func() { channel <- 1 }() // send 1 to the channel
	go func() { channel <- 2 }() // send 2 to the channel
	go func() { channel <- 3 }() // send 3 to the channel
	// after it sends to the channel, it starts waiting

	// receive from the channel (nondeterministic)
	first := <-channel
	second := <-channel
	third := <-channel
	fmt.Println(first, second, third)

	// Buffered channel
	channel1 := make(chan int, 2) // 2 is the buffer size
	channel1 <- 1
	channel1 <- 2
	// until buffer is full, it doesn't wait
	go func() { channel1 <- 3 }()
	// after buffer is full, it starts waiting (the async function blocks because the buffer is full 3 > 2 aka buffer size)

	first1 := <-channel1
	second1 := <-channel1
	third1 := <-channel1

	fmt.Println(first1, second1, third1)

	// control channel
	one := make(chan int)
	two := make(chan int)

	//for {
	//	select {
	//	// blocked when no data is available
	//	case o := <-one:
	//		fmt.Println("one", o)
	//	case t := <-two:
	//		fmt.Println("two", t)
	//	// first two blocked if no data, so we go ot default
	//	default:
	//		fmt.Println("no data")
	//		time.Sleep(50 * time.Millisecond)
	//	}
	//
	//}

	for {
		select {
		// blocked when no data is available
		case o := <-one:
			fmt.Println("one", o)
		case t := <-two:
			fmt.Println("two", t)
		// first two blocked if no data, so we go ot default
		case <-time.After(500 * time.Millisecond):
			fmt.Println("Timeout: No data received")
			return
		}
	}

}
