package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := getInput()
	fmt.Println(calculatePresentsA(input))
	fmt.Println(calculatePresentsB(input))
}

func calculatePresentsA(input string) int {
	x := 0
	y := 0
	startLocation := []int{x, y}
	allLocations := [][]int{startLocation}

	for _, move := range strings.Split(input, "") {
		if move == "<" {
			x -= 1
		}
		if move == ">" {
			x += 1
		}
		if move == "^" {
			y += 1
		}
		if move == "v" {
			y += -1
		}
		newLocation := []int{x, y}
		if !contains(allLocations, newLocation) {
			allLocations = append(allLocations, newLocation)
		}
	}
	return len(allLocations)
}

func calculatePresentsB(input string) int {
	x := 0
	y := 0
	x_robo := 0
	y_robo := 0
	startLocation := []int{x, y}
	allLocations := [][]int{startLocation}

	for i, move := range strings.Split(input, "") {
		if move == "<" {
			if i%2 == 0 {
				x -= 1
			} else {
				x_robo -= 1
			}
		}
		if move == ">" {
			if i%2 == 0 {
				x += 1
			} else {
				x_robo += 1
			}
		}
		if move == "^" {
			if i%2 == 0 {
				y -= 1
			} else {
				y_robo -= 1
			}
		}
		if move == "v" {
			if i%2 == 0 {
				y += 1
			} else {
				y_robo += 1
			}
		}
		newLocation := []int{x, y}
		newLocation_robo := []int{x_robo, y_robo}
		if !contains(allLocations, newLocation) {
			allLocations = append(allLocations, newLocation)
		}
		if !contains(allLocations, newLocation_robo) {
			allLocations = append(allLocations, newLocation_robo)
		}
	}
	return len(allLocations)
}

func getInput() string {
	bs, err := os.ReadFile("input")
	if err != nil {
		fmt.Println("Error reading input file: ", err)
		os.Exit(1)
	}
	return string(bs)
}

func contains(locations [][]int, location []int) bool {
	for _, loc := range locations {
		if location[0] == loc[0] && location[1] == loc[1] {
			return true
		}
	}
	return false
}
