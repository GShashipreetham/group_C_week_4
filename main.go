package main

import "fmt"

func main() {
  fmt.Println("Welcome to Group C's Week 4!")

func reverseNumber(n int) int{
  reversed := 0
  for n != 0{
    remainder := n%10
    reversed = reversed*10 + remainder

  }
