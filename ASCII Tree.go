package main

import "fmt"

func Tree(height int) {
	for i := 0; i < height; i++ {
		for j := height - i - 1; j > 0; j-- {
			fmt.Print(" ")
		}
		for l := 0; l < (7*i + 1); l++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := 0; i < height-1; i++ {
		fmt.Print(" ")
	}
	fmt.Println("*")
}
