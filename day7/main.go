package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type circuit map[string]int
type instruction []string

func main() {
	input := strings.Split(getInput(), "\n")
	instructions := []instruction{}
	circuit := make(circuit)
	for _, line := range input {
		instructions = append(instructions, strings.Split(line, " -> "))
	}
	secondRun := false
	for {
		for _, instruction := range instructions {
			_, found := circuit[instruction[1]]
			if !found {
				operands, err := parseOperands(getOperands(instruction[0]), circuit)
				if err == nil {
					circuit[instruction[1]] = processInstruction(determineOperator(instruction[0]), operands)
					if instruction[1] == "a" || len(circuit) == len(input) {
						fmt.Println(circuit["a"])
						if secondRun {
							os.Exit(0)
						}
						newB := circuit["a"]
						for i := range circuit {
							delete(circuit, i)
						}
						circuit["b"] = newB
						secondRun = true
					}
				}
			}
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

func determineOperator(leftSide string) string {
	if strings.Contains(leftSide, "AND") {
		return "AND"
	}
	if strings.Contains(leftSide, "OR") {
		return "OR"
	}
	if strings.Contains(leftSide, "LSHIFT") {
		return "LSHIFT"
	}
	if strings.Contains(leftSide, "RSHIFT") {
		return "RSHIFT"
	}
	if strings.Contains(leftSide, "NOT") {
		return "NOT"
	}
	return "DIRECT"
}

func getOperands(leftSide string) []string {
	operator := determineOperator(leftSide)
	leftSideSplit := strings.Split(leftSide, " ")
	if operator == "AND" || operator == "OR" || operator == "LSHIFT" || operator == "RSHIFT" {
		return []string{leftSideSplit[0], leftSideSplit[2]}
	}
	if operator == "NOT" {
		return []string{leftSideSplit[1]}
	}
	return []string{leftSideSplit[0]}
}

func parseOperands(operands []string, c circuit) ([]int, error) {
	intOps := []int{}
	for _, op := range operands {
		intOp, err := strconv.Atoi(op)
		if err != nil {
			foundInt, found := c[op]
			if !found {
				return []int{}, err
			}
			intOp = foundInt
		}
		intOps = append(intOps, intOp)
	}
	return intOps, nil
}

func processInstruction(operator string, operands []int) int {
	if operator == "AND" {
		return bitwiseAnd(operands)
	}
	if operator == "OR" {
		return bitwiseOr(operands)
	}
	if operator == "NOT" {
		return bitwiseNot(operands[0])
	}
	if operator == "LSHIFT" {
		return bitwiseLeftShift(operands)
	}
	if operator == "RSHIFT" {
		return bitwiseRightShift(operands)
	}
	return operands[0]
}

func bitwiseAnd(operands []int) int {
	binary1 := convertIntTo16Bit(operands[0])
	binary2 := convertIntTo16Bit(operands[1])
	binaryResult := []rune("0000000000000000")

	for i := range binary1 {
		if binary1[i] == '1' && binary2[i] == '1' {
			binaryResult[i] = '1'
		}
	}
	decimal, _ := strconv.ParseUint(string(binaryResult), 2, 16)
	return int(decimal)
}

func bitwiseOr(operands []int) int {
	binary1 := convertIntTo16Bit(operands[0])
	binary2 := convertIntTo16Bit(operands[1])
	binaryResult := []rune("0000000000000000")

	for i := range binary1 {
		if binary1[i] == '1' || binary2[i] == '1' {
			binaryResult[i] = '1'
		}
	}
	decimal, _ := strconv.ParseUint(string(binaryResult), 2, 16)
	return int(decimal)
}

func bitwiseNot(operand int) int {
	binary := convertIntTo16Bit(operand)

	binaryResult := []rune("0000000000000000")

	for i := range binary {
		if binary[i] == '0' {
			binaryResult[i] = '1'
		}
	}
	decimal, _ := strconv.ParseUint(string(binaryResult), 2, 16)
	return int(decimal)
}

func bitwiseLeftShift(operands []int) int {
	binary := convertIntTo16Bit(operands[0])
	shifts := operands[1]

	binaryResult := ""
	for i := 0; i < 16-shifts; i++ {
		binaryResult += string(binary[i+shifts])
	}
	for i := 0; i < shifts; i++ {
		binaryResult += "0"
	}
	decimal, _ := strconv.ParseUint(string(binaryResult), 2, 16)
	return int(decimal)
}

func bitwiseRightShift(operands []int) int {
	binary := convertIntTo16Bit(operands[0])
	shifts := operands[1]

	binaryResult := ""
	for i := 0; i < shifts; i++ {
		binaryResult += "0"
	}
	for i := shifts; i < 16; i++ {
		binaryResult += string(binary[i-shifts])
	}

	decimal, _ := strconv.ParseUint(string(binaryResult), 2, 16)
	return int(decimal)
}

func convertIntTo16Bit(i int) string {
	binary := strconv.FormatInt(int64(i), 2)
	leadingZeros := ""
	for i := 0; i < 16-len(binary); i++ {
		leadingZeros += "0"
	}
	return leadingZeros + binary
}
