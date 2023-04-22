package main

import "fmt"

func main() {
	fmt.Println("Hi")
	coord := Coordinate{1, 2}
	shiftBy(3, 4, &coord)
	// shiftBy with receiver function
	coord.shiftBy(3, 4)
	fmt.Println(&coord)
	dist := coord.shiftDist(3, 4)
	fmt.Println(dist)
	fmt.Println(coord)

}

// Regular functions
type Coordinate struct {
	X int
	Y int
}

func shiftBy(x, y int, coord *Coordinate) {
	coord.X += x
	coord.Y += y
}

// Receiver functions
func (coord *Coordinate) shiftBy(x, y int) {
	coord.X += x
	coord.Y += y
}

// Value receiver - receives a copy of the struct
func (coord Coordinate) shiftDist(x, y int) Coordinate {
	coord.X += x
	coord.Y += y
	return coord
}
