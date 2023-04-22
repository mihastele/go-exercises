package msg

import "fmt"

func Exciting(msg string) {
	// you need to go from the root package go.mod file to the package you want to import
	fmt.Printf("%v!", msg)
}
