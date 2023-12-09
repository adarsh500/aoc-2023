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
			count++
		}
	}

	return int(math.Pow(2, float64(count)-1))
}

var frequency = map[int]int{}

func getScratchCardsCount(index int, line string) {
	frequency[index]++
	counter := map[int]int{}
	count := 0
	repetition := frequency[index]
	cards := strings.Split(strings.Split(line, ":")[1], "|")
	winningCards, regularCards := strings.Split(strings.Trim(cards[0], " "), " "), strings.Split(strings.Trim(cards[1], " "), " ")

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
			count++
		}
	}

	fmt.Println("c", count, repetition)
	for count+index > index {
		frequency[count+index] += repetition
		count--
	}

}

func main() {
	points := 0
	data, err := os.ReadFile("./input.txt")
	check(err)
	lines := strings.Split(string(data), "\n")

	for i, line := range lines {
		fmt.Println(frequency)
		getScratchCardsCount(i+1, line)
	}

	// fmt.Println(frequency)

	for _, value := range frequency {
		points += value
	}

	fmt.Println("result", points)
}
