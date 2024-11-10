package main

import "fmt"

func main() {
	value := 10

	fmt.Println(add(pow(value), 1))
}

func pow(a int) int {
	return a * a
}

func add(a int, b int) int {
	return a + b
}
