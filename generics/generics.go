package main

func name[T any, U int | string](v T, u U) {
	println(v, u)
}

func isEqual[T comparable](a, b T) bool {
	return a == b
}

func main() {
	name[int, int](1, 2)
	name[string, int]("hello", 1)
	name[string, string]("hello", "world")

	println(isEqual[int](1, 2))
	println(isEqual[int](1, 1))
	println(isEqual[string]("hello", "world"))
	println(isEqual[string]("hello", "hello"))
}
