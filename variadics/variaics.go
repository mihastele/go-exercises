package main

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	println(sum(1, 2, 3, 4, 5))
	a := []int{1, 2, 3, 4, 5}
	b := []int{6, 7, 8, 9, 10}
	all := append(a, b...)
	println(sum(all...))
}
