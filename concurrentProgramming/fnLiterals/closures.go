package main

import "fmt"

func helloWorld() {
	fmt.Printf("Hello, ")
	// anoymous function
	world := func() {
		fmt.Printf("world!\n")
	}

	world()
	world()

	// Closures
	// A closure is a function value that references variables from outside its body.
	// The function may access and assign to the referenced variables; in this sense
	// the function is "bound" to the variables.
	// For example, the adder function returns a closure. Each closure is bound to its
	// own sum variable.
	outside := 1
	worldC := func() {
		outside++
	}

	// Higher order function
	world2 := func() func() {
		return func() {
			fmt.Printf("world!\n")
		}
	}

	worldC()
	worldC()
	worldC()
	world2()()

	fmt.Println(outside)

}

func main() {
	helloWorld()
}
