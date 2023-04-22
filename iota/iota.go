package main

import "fmt"

type Direction byte

const (
	North Direction = iota
	East
	South
	West
)

func main() {
	//const (
	//	Online      = 0
	//	Offline     = 1
	//	Maintenance = 2
	//	Retired     = 3
	//)

	// iota is a special constant that can be used to create a set of related but distinct constants
	const (
		Online = iota
		Offline
		Maintenance
		Retired
	)
	fmt.Println(North)

}

//func (d Direction) String() string {
//	switch d {
//	case North:
//		return "North"
//	case East:
//		return "East"
//	case South:
//		return "South"
//	case West:
//		return "West"
//	default:
//		return "Unknown"
//	}
//}

// shorter version with slicing
func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}
