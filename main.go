package main

import "fmt"

func main() {

	fmt.Println("Welcome to Group C's Week 4!")

	var num int

	fmt.Print("Enter a number to check if it's a Happy Number: ")
	fmt.Scanln(&num)

	if isHappyNumber(num) {
		fmt.Printf("%d is a Happy Number.\n", num)
	} else {
		fmt.Printf("%d is not a Happy Number.\n", num)
	}

	var s int
	fmt.Print("Enter the number of rows for the butterfly pattern: ")
	fmt.Scan(&s)
	cloud2003(s)
	var n int
	fmt.Print("Enter the number you want to reverse: ")
	fmt.Scan(&n)
	print("reverse of the number is:\n", numberreverse(n))

	var height int
	fmt.Println("Enter the height of the tree: ")
	fmt.Scan(&height)

	if height <= 0 {
		fmt.Println("Height should be a positive integer.")
		return
	}

	Tree(height)

	
	var input string
	fmt.Print("Enter the word:")
	fmt.Scan(&input)
	print("Rainbow word is: ",)
	RainbowText(input)
	
	fmt.Println() 


	myCar := Car{
		Vehicle: Vehicle{make: "Toyota", model: "Corolla"},
		seats:   5,
	}
	myCar.Start()
	myCar.Drive()

	myEV := ElectricVehicle{
		Vehicle:        Vehicle{make: "Tesla", model: "Model S"},
		batteryCapacity: 100,
	}
	myEV.Start()
	myEV.Drive()
	myEV.Charge()

}
