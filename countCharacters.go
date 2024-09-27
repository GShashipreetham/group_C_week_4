package main

func countCharacters(s string) map[rune]int {
	countMap := make(map[rune]int)
	for _, char := range s{
		countMap[char]++
	}
	return countMap

}