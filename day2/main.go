package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := getInput()
	wrappingPaper := []int{}
	ribbon := []int{}
	for _, line := range strings.Split(input, "\n") {
		splitLine := strings.Split(line, "x")
		l, err := strconv.Atoi(splitLine[0])
		w, err := strconv.Atoi(splitLine[1])
		h, err := strconv.Atoi(splitLine[2])

		smallest_1 := int(math.Min(float64(l), float64(w)))
		smallest_2 := int(math.Min(math.Max(float64(l), float64(w)), float64(h)))

		if err != nil {
			fmt.Println("Error parsing string to int: ", err)
			os.Exit(1)
		}
		wrappingPaper = append(wrappingPaper, 2*l*w+2*w*h+2*h*l+smallest_1*smallest_2)
		ribbon = append(ribbon, 2*smallest_1+2*smallest_2+l*w*h)
	}
	sumWrappingPaper := 0
	sumRibbon := 0
	for i, wrappingPaperSize := range wrappingPaper {
		sumWrappingPaper += wrappingPaperSize
		sumRibbon += ribbon[i]
	}
	fmt.Println(sumWrappingPaper)
	fmt.Println(sumRibbon)
}

func getInput() string {
	bs, err := os.ReadFile("input")
	if err != nil {
		fmt.Println("Error reading input file: ", err)
		os.Exit(1)
	}
	return string(bs)
}
