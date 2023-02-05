package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := getInput()
	fmt.Println("The final floor:", determineFloor(input))

	for i := range input {
		if determineFloor(input[:i+1]) == -1 {
			fmt.Println("The first instance basement:", i+1)
			os.Exit(0)
		}
	}
}

func getInput() string {
	bs, err := os.ReadFile("input")
	if err != nil {
		fmt.Println("Error reading input file: ", err)
		os.Exit(1)
	}
	return string(bs)
}

func determineFloor(instructions string) int {
	open := strings.Count(instructions, "(")
	close := strings.Count(instructions, ")")
	return open - close
}
