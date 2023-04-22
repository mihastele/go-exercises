package msg

// you need to go from the root package go.mod file to the package you want to import
import "hi/packages/display"

func Hi() {

	display.Display("Hi!")
}
