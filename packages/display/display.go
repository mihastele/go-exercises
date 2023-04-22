package display

import "fmt"

// since it starts with a capital letter, it is exported
func Display(msg string) {
	fmt.Println(msg)
}

// since it starts with a lowercase letter, it is not exported and can only be used within this package
func hello(msg string) {
	fmt.Println("Hello", msg)
}
