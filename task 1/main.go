package main

import "fmt"
func main() {
	fmt.Println("Hello, World!")


	fmt.Println("Welcome to Task 1")
	sum := sumSlices([]int{1, 2, 3, 4, 5})
	fmt.Println("Sum of slices:", sum)
}

func sumSlices(slices []int) int {
	sum := 0
	for _, slice := range slices {
		sum += slice
	}
	return sum
}