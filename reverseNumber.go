package main

func numberreverse(num int) int {
	reversed := 0
	for num > 0 {
		reversed = reversed*10 + num%10
		num /= 10
	}
	return reversed
}
