package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func isValidConfiguration(line string) int {
	bag := map[string]int{"red": 12, "green": 13, "blue": 14}

	game := strings.Split(line, ": ")
	_, id, _ := strings.Cut(game[0], "Game ")
	sets := strings.Split(game[1], "; ")

	for _, set := range sets {
		count := make(map[string]int)

		draw := strings.Split(set, ", ")
		for _, ball := range draw {
			data := strings.Split(ball, " ")
			value, color := data[0], data[1]
			valueInt, _ := strconv.Atoi(value)
			count[color] += valueInt
		}

		for key, value := range count {
			if value > bag[key] {
				return 0
			}
		}
	}

	var idInt, _ = strconv.Atoi(id)
	return idInt
}

func getPowerOfCubes(line string) int {
	game := strings.Split(line, ": ")
	sets := strings.Split(game[1], "; ")
	count := make(map[string]int)

	for _, set := range sets {
		draw := strings.Split(set, ", ")
		for _, ball := range draw {
			data := strings.Split(ball, " ")
			value, color := data[0], data[1]
			valueInt, _ := strconv.Atoi(value)
			count[color] = max(valueInt, count[color])
		}
	}

	result := 1
	for _, value := range count {
		result *= value
	}

	return result
}

func main() {
	result := 0
	data, err := os.ReadFile("./input.txt")
	check(err)
	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		result += getPowerOfCubes(line)
	}

	fmt.Println("result", result)
}
