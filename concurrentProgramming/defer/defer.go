package main

import (
	"fmt"
	"os"
)

func mihaDeferred() {
	fmt.Println("Hi, I'm Miha")
}

func deferSample() {
	// Defer
	// A defer statement defers the execution of a function until the surrounding function returns.
	// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
	// Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
	// For example, this function prints "deferred" after the function returns.
	// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
	defer func() {
		println("deferred")
	}()

	defer mihaDeferred()

	println("done")
}

func main() {
	deferSample()

	file, err := os.Open("concurrentProgramming/defer/hii.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Oh no!")
		}
	}(file)

	buffer := make([]byte, 0, 30)
	bytes, err := file.Read(buffer)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Bytes read: %d \t Data: %s", bytes, buffer)

}
