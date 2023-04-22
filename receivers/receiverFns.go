package main

import "fmt"

func main() {
	fmt.Println("Hi")
	coord := Coordinate{1, 2}
	coord.shiftBy(3, 4)
	fmt.Println(&coord)

}

// Regular functions
type Coordinate struct {
	X int
	Y int
}

//
//func shiftBy(x, y int, coord *Coordinate) {
//	coord.X += x
//	coord.Y += y
//}

// Receiver functions
func (coord *Coordinate) shiftBy(x, y int) {
	coord.X += x
	coord.Y += y
}
