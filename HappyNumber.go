package main

func isHappyNumber(n int) bool {
    seen := map[int]bool{}
    for n != 1 && !seen[n] {
        seen[n] = true
        n = sumOfSquaresOfDigits(n)
    }
    return n == 1
}


func sumOfSquaresOfDigits(n int) int {
    sum := 0
    for n > 0 {
        digit := n % 10
        sum += digit * digit
        n /= 10
    }
    return sum
}

