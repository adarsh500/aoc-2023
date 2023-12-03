// unoptimized
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var directions = [8][2]int{{0, 1}, {1, 0}, {1, -1}, {-1, -1}, {0, -1}, {-1, 0}, {1, 1}, {-1, 1}}
var counted = map[int][]string{}

func isNumber(value string) bool {
	_, err := strconv.Atoi(value)
	return err == nil
}

func isSymbol(char rune) bool {
	value := string(char)
	return value != "." && !isNumber(value)
}

func getAdjacentNumber(i int, j int, line string) int {
	value := string(line[j])
	start, end := j, j

	for index := j + 1; index < len(line); index++ {
		if isNumber(string(line[index])) {
			value += string(line[index])
			end = index
		} else {
			break
		}
	}

	for index := j - 1; index >= 0; index-- {
		if isNumber(string(line[index])) {
			value = string(line[index]) + value
			start = index
		} else {
			break
		}
	}

	if slices.Contains(counted[i], string(start+end)) {
		return -1
	} else {
		newSlice := append(counted[i], string(start+end))
		counted[i] = newSlice
	}

	valueInt, _ := strconv.Atoi(value)
	return valueInt
}

func getAllNumbers(i int, j int, lines []string) int {
	count, gearRatio := 0, 1
	isStar := string(lines[i][j]) == "*"
	if !isStar {
		return 0
	}

	for _, direction := range directions {
		dx, dy := i+direction[0], j+direction[1]

		if isNumber(string(lines[dx][dy])) {
			adjNumber := getAdjacentNumber(dx, dy, lines[dx])
			if adjNumber > 0 {
				gearRatio *= adjNumber
				count++
			}
		}
	}

	if count != 2 {
		return 0
	}

	return gearRatio
}

func main() {
	result := 0
	data, err := os.ReadFile("./input.txt")
	check(err)
	lines := strings.Split(string(data), "\n")

	for i, line := range lines {
		for j, char := range line {
			if isSymbol(char) {
				result += getAllNumbers(i, j, lines)
			}
		}
	}

	fmt.Println("result", result)
}
