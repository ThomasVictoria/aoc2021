package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	aoc "github.com/ThomasVictoria/advent_of_code_2021"
)

type Position struct {
	Actual   int
	Left     int
	Right    int
	Bottom   int
	Top      int
	IsLowest bool
}

func (p *Position) Init(array []string, currentLine int, currentIndex int) {

	p.Actual = convertStringToInt(
		fmt.Sprintf("%c", array[currentLine][currentIndex]),
	)

	if currentIndex-1 >= 0 {
		p.Left = convertStringToInt(
			fmt.Sprintf("%c", array[currentLine][currentIndex-1]),
		)
	} else {
		p.Left = 99
	}

	if len(array[currentLine]) > currentIndex+1 {
		p.Right = convertStringToInt(
			fmt.Sprintf("%c", array[currentLine][currentIndex+1]),
		)
	} else {
		p.Right = 99
	}

	if currentLine-1 >= 0 {
		p.Top = convertStringToInt(
			fmt.Sprintf("%c", array[currentLine-1][currentIndex]),
		)
	} else {
		p.Top = 99
	}

	if len(array) > currentLine+1 {
		p.Bottom = convertStringToInt(
			fmt.Sprintf("%c", array[currentLine+1][currentIndex]),
		)
	} else {
		p.Bottom = 99
	}
}

func (p *Position) isLowest() bool {

	if p.Actual < p.Left &&
		p.Actual < p.Right &&
		p.Actual < p.Top &&
		p.Actual < p.Bottom {
		p.IsLowest = true
		return true
	}

	p.IsLowest = false
	return false
}

func convertStringToInt(in string) int {
	res, err := strconv.Atoi(in)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func main() {

	var sumOfRisks int
	var bassinMap [][]Position

	data := strings.Split(aoc.DecodeInput("9"), "\n")

	for lineIndex, line := range data {

		bassinMap = append(bassinMap, make([]Position, 100))

		for index, _ := range line {

			position := Position{}
			position.Init(data, lineIndex, index)

			bassinMap[lineIndex][index] = position

			if position.isLowest() {
				sumOfRisks += (position.Actual + 1)
			}
		}
	}

	fmt.Println(sumOfRisks)

	for lineIndex, line := range bassinMap {
		for index, position := range line {
			if position.IsLowest {
				propagate(bassinMap, lineIndex, index)
			}
		}
	}

}

func propagate(bassins [][]Position, lineIndex int, index int) bool {

	if lineIndex-1 >= 0 && bassins[lineIndex-1][index].IsLowest {
		return propagate(bassins, lineIndex-1, index)
	}
	// Unfinished
	return true
}
