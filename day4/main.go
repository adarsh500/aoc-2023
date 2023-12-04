package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func computePoints(line string) int {
	counter := map[int]int{}
	count := 0
	cards := strings.Split(strings.Split(line, ":")[1], "|")
	winningCards, regularCards := strings.Split(strings.Trim(cards[0], " "), " "), strings.Split(strings.Trim(cards[1], " "), " ")
	fmt.Println(winningCards)
	fmt.Println(regularCards)
	for _, cards := range winningCards {
		cardNum, _ := strconv.Atoi(strings.Trim(cards, " "))
		counter[cardNum] = 1
	}

	for _, cards := range regularCards {
		cardNum, _ := strconv.Atoi(strings.Trim(cards, " "))
		if cardNum == 0 || cards == "" {
			continue
		}

		if counter[cardNum] == 1 {
			fmt.Print(cardNum, ",")
			count++
		}
	}
	fmt.Println()
	// fmt.Println(counter)

	return int(math.Pow(2, float64(count)-1))
}

func main() {
	points := 0
	data, err := os.ReadFile("./input.txt")
	check(err)
	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		points += computePoints(line)
	}

	fmt.Println("result", points)
}
