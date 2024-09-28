package main

import "fmt"

func main() {
<<<<<<< HEAD

  fmt.Println("Welcome to Group C's Week 4!")
  
  fmt.Println("factorial is:" ,)
  
  
    var num int

    // Ask the user to input a number
    fmt.Print("Enter a number to check if it's a Happy Number: ")
    fmt.Scanln(&num)

    // Check if the entered number is a Happy Number
    if isHappyNumber(num) {
        fmt.Printf("%d is a Happy Number.\n", num)
    } else {
        fmt.Printf("%d is not a Happy Number.\n", num)
    }
}








    

=======
	var s int
	fmt.Print("Enter the number of rows for the butterfly pattern: ")
	fmt.Scan(&s)
	cloud2003(s)
}
>>>>>>> 20b75a599b85aade3db01954b7b8c70dc8d0ffad
