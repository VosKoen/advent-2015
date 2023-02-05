package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type location = [2]int
type grid map[location]bool
type gridPart2 map[location]int

func main() {
	grid := make(grid)
	gridPart2 := make(gridPart2)
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			grid[location{i, j}] = false
			gridPart2[location{i, j}] = 0
		}
	}
	input := strings.Split(getInput(), "\n")
	for _, line := range input {
		splitLine := strings.Split(line, " ")
		if len(splitLine) == 4 {
			grid.toggle(convertStringLocationToLocation(splitLine[1]), convertStringLocationToLocation(splitLine[3]))
			gridPart2.increaseBrightness(convertStringLocationToLocation(splitLine[1]), convertStringLocationToLocation(splitLine[3]), 2)
		}
		if splitLine[1] == "off" {
			grid.turnOff(convertStringLocationToLocation(splitLine[2]), convertStringLocationToLocation(splitLine[4]))
			gridPart2.decreaseBrightness(convertStringLocationToLocation(splitLine[2]), convertStringLocationToLocation(splitLine[4]))
		}
		if splitLine[1] == "on" {
			grid.turnOn(convertStringLocationToLocation(splitLine[2]), convertStringLocationToLocation(splitLine[4]))
			gridPart2.increaseBrightness(convertStringLocationToLocation(splitLine[2]), convertStringLocationToLocation(splitLine[4]), 1)
		}
	}
	fmt.Println(grid.countLights())
	fmt.Println(gridPart2.countBrightness())
}

func getInput() string {
	bs, err := os.ReadFile("input")
	if err != nil {
		fmt.Println("Error reading input file: ", err)
		os.Exit(1)
	}
	return string(bs)
}

func convertStringLocationToLocation(s string) location {
	splitLocation := strings.Split(s, ",")
	firstCoordinate, _ := strconv.Atoi(splitLocation[0])
	secondCoordinate, _ := strconv.Atoi(splitLocation[1])
	return location{firstCoordinate, secondCoordinate}
}

func (g grid) turnOn(start location, end location) {
	for i := start[0]; i <= end[0]; i++ {
		for j := start[1]; j <= end[1]; j++ {
			g[location{i, j}] = true
		}
	}
}

func (g grid) turnOff(start location, end location) {
	for i := start[0]; i <= end[0]; i++ {
		for j := start[1]; j <= end[1]; j++ {
			g[location{i, j}] = false
		}
	}
}

func (g grid) toggle(start location, end location) {
	for i := start[0]; i <= end[0]; i++ {
		for j := start[1]; j <= end[1]; j++ {
			g[location{i, j}] = !g[location{i, j}]
		}
	}
}

func (g grid) countLights() int {
	count := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if g[location{i, j}] == true {
				count++
			}
		}
	}
	return count
}

func (g gridPart2) countBrightness() int {
	count := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			count += g[location{i, j}]
		}
	}
	return count
}

func (g gridPart2) increaseBrightness(start location, end location, increase int) {
	for i := start[0]; i <= end[0]; i++ {
		for j := start[1]; j <= end[1]; j++ {
			g[location{i, j}] += increase
		}
	}
}

func (g gridPart2) decreaseBrightness(start location, end location) {
	for i := start[0]; i <= end[0]; i++ {
		for j := start[1]; j <= end[1]; j++ {
			if g[location{i, j}] > 0 {
				g[location{i, j}] -= 1
			}
		}
	}
}
