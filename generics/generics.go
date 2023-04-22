package main

func name[T any, U int | string](v T) {
	println(v)
}

func isEqual[T comparable](a, b T) bool {
	return a == b
}

func main() {
	name[int, int](1)
	name[int, string](2)
	name[string, int]("hello")
	name[string, string]("world")

	println(isEqual[int](1, 2))
	println(isEqual[int](1, 1))
	println(isEqual[string]("hello", "world"))
	println(isEqual[string]("hello", "hello"))
}
