package main

import "fmt"

func double(x int) int {
	return x + x
}

func greet() {
	fmt.Println("Hello World!")
}

type Sample struct {
	field string
	a     int
	b     int
}

func dereference(x *int) {
	*x = *x * 2
	*x += 1
}

func main() {
	greet()

	dozen := double(6)
	fmt.Println(dozen + 1)

	for i := 0; i < 10; i++ {
		data := Sample{"hello", i, i * i}
		fmt.Println(data.b)
	}

	var myArray [10]int
	myArray1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	myArray2 := [...]int{1, 2, 3, 4}
	myArray3 := [4]int{1, 3}
	myArray4 := [4]int{1: 3, 3: 5}
	myArray[0] = 4

	j := 0
	for j < cap(myArray)/2 {
		myArray[j+1] = j + 2
		j++
	}

	slice1 := myArray[1:3]
	slice := []int{1, 2, 3}
	slice = append(slice, 4)
	slice = append(slice, 5, 6, 20)
	slice = append(slice, slice1...)

	fmt.Println(myArray, myArray1, myArray2, myArray3, myArray4, slice)

	// we can only use range to iterate through characters in a string because emoji use more than a byte for example
	sliceS := []string{"miha", "iha", "ha", "ahim"}
	for i, element := range sliceS {
		fmt.Println(i, element, ":")
		for _, ch := range element {
			fmt.Printf(" %q\n", ch)
		}
	}

	mymap1 := make(map[string]int)
	mymap1["one"] = 1
	mymap2 := map[string]int{"one": 5, "two": 2}
	delete(mymap1, "one")
	fmt.Println(mymap1, mymap2)

	price, found := mymap2["one"]
	fmt.Println(price, found)
	price2, found := mymap1["one"]
	fmt.Println(price2, found)

	for key, value := range mymap2 {
		fmt.Println(key, value)
	}

	value := 10
	var valuePtr *int = &value
	fmt.Println(valuePtr, *valuePtr)

	dereference(valuePtr)
	fmt.Println(value)
}
