package main

import "fmt"

func cloud2003(s int) {

	for i := 1; i <= s; i++ {

		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}

		for j := 1; j <= 2*(s-i); j++ {
			fmt.Print(" ")
		}

		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}

	for i := s; i >= 1; i-- {

		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}

		for j := 1; j <= 2*(s-i); j++ {
			fmt.Print(" ")
		}

		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}
