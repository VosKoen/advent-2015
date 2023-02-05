package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := strings.Split(getInput(), "\n")
	count := 0
	count2 := 0
	for _, s := range input {
		if isNice(s) {
			count += 1
		}
		if isNiceSecondPart(s) {
			count2 += 1
		}
	}
	fmt.Println(count)
	fmt.Println(count2)
}

func getInput() string {
	bs, err := os.ReadFile("input")
	if err != nil {
		fmt.Println("Error reading input file: ", err)
		os.Exit(1)
	}
	return string(bs)
}

func isNice(s string) bool {
	return countVowels(s) >= 3 && hasDoubleLetters(s) && !containsForbiddenStrings(s)
}

func isNiceSecondPart(s string) bool {
	return containsTwoLetterDuplicate(s) && repeatsWithOneLetterBetween(s)
}

func countVowels(s string) int {
	count := 0
	for _, char := range s {
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			count += 1
		}
	}
	return count
}

func hasDoubleLetters(s string) bool {
	for i, char := range s {
		if i > 0 && rune(s[i-1]) == char {
			return true
		}
	}
	return false
}

func containsForbiddenStrings(s string) bool {
	return strings.Contains(s, "ab") || strings.Contains(s, "cd") || strings.Contains(s, "pq") || strings.Contains(s, "xy")
}

func containsTwoLetterDuplicate(s string) bool {
	for i := 0; i < len(s)-3; i++ {
		if strings.Contains(s[i+2:], s[i:i+2]) {
			return true
		}
	}
	return false
}

func repeatsWithOneLetterBetween(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}
