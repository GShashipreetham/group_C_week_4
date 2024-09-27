package main

import "fmt"

func main() {
	fmt.Println("Welcome to Group C's Week 4!")

	arr := []int{11, 15, 19, 25, 30, 40}
	fmt.Println("Unsorted array:", arr)

	BubbleSort(arr)

	fmt.Println("Sorted array: ", arr)
}
