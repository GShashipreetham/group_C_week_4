package main

import "fmt"

var colors = []string{
    "\033[31m", "\033[33m","\033[32m","\033[36m","\033[34m","\033[35m","\033[0m", 
}

func RainbowText(input string) {
    for i, char := range input {
        color := colors[i%len(colors)]
        fmt.Print(color, string(char))
    }
    fmt.Print(colors[len(colors)-1])
}



