package main

import "fmt"

var colors = []string{
    "\033[38;5;135m", "\033[38;5;93m","\033[34m","\033[32m","\033[33m","\033[38;5;214m","\033[31m", 
}

func RainbowText(input string) {
    for i, char := range input {
        color := colors[i%len(colors)]
        fmt.Print(color, string(char))
    }
    fmt.Print("\033[0m")

}



