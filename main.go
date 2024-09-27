package main

import "fmt"

func main() {
  fmt.Println("Welcome to Group C's Week 4!")

func revesedNumber(){
  number := 12345
  reversed := 0
  for number != 0{
    remainder := number%10
    reversed = reversed*10 + remainder
    number /=10
    }
  fmt.Print("Reversednumber.%d\n", reversed) }


func countCharacters(s string) map[rune]int {
      countMap := make(map[rune]int)
      for _, char := range s {
        countMap[char]++
      }
      return countMap
  }

    
  }
