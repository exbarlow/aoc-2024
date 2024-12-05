package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Coord struct {
	row, col int
}

const search string = "XMAS"

func main() {
	b, err := os.ReadFile("day4.txt")

	if err != nil {
		log.Fatal(err)
	}

	inputStr := string(b)
	lines := strings.Split(inputStr, "\n")

	var letters [][]byte;
	
	for _, line := range lines {
		letters = append(letters, []byte(line))
	}
	part1 := 0

	for row := range letters {
		for col := range letters[row] {
			part1 += traverse(letters, Coord{row, col}, Coord{-1, -1}, 0)
		}
	}

	part2 := 0

	for row := 0; row < len(letters) - 2; row++ {
		for col := 0; col < len(letters[0]) - 2; col++ {
			cl := ConcatLetters(letters, row, col)
			if cl == "MSAMS" || cl == "MMASS" || cl == "SSAMM" || cl == "SMASM" {
				part2++
			}
		}
	}

	fmt.Printf("part1: %d\n", part1)
	fmt.Printf("part2: %d\n", part2)
}

func (c *Coord) Add(a Coord) {
	c.col += a.col
	c.row += a.row
}

func sumCoords(c, a Coord) Coord {
	return Coord{a.row + c.row, a.col + c.col}
}

func (c Coord) Diff(c2 Coord) Coord {
	return Coord{c.row - c2.row, c.col - c2.col}
}

func ConcatLetters(letters [][]byte, row, col int) string {
	var b []byte

	b = append(b, letters[row][col])
	b = append(b, letters[row][col+2])
	b = append(b, letters[row+1][col+1])
	b = append(b, letters[row+2][col])
	b = append(b, letters[row+2][col+2])

	return string(b)
}

func traverse(letters [][]byte, currCoord, prevCoord Coord, searchIdx int) int {
	if currCoord.row < 0 || currCoord.col < 0 || currCoord.row >= len(letters) || currCoord.col >= len(letters[0]) {
		return 0
	}


	currLetter := letters[currCoord.row][currCoord.col]

	if currLetter != search[searchIdx] {
		return 0
	}
	// fmt.Println(currCoord)
	// fmt.Println(prevCoord)
	// fmt.Println(string(currLetter))

	if searchIdx == len(search) - 1 {
		// fmt.Println("**** FOUND ****")
		return 1
	}

	sum := 0

	if prevCoord == (Coord{-1, -1}) || currCoord.Diff(prevCoord) == (Coord{-1, -1}) {
		// fmt.Println("going up and left")
		// fmt.Println(currCoord)
		// fmt.Println(prevCoord)
		sum += traverse(letters, sumCoords(currCoord, Coord{-1, -1}), currCoord, searchIdx + 1)
	}
	if prevCoord == (Coord{-1, -1}) || currCoord.Diff(prevCoord) == (Coord{0, -1}) {
		sum += traverse(letters, sumCoords(currCoord, Coord{0, -1}), currCoord, searchIdx + 1)
	}
	if prevCoord == (Coord{-1, -1}) || currCoord.Diff(prevCoord) == (Coord{1, -1}) {
		sum += traverse(letters, sumCoords(currCoord, Coord{1, -1}), currCoord, searchIdx + 1)
	}
	if prevCoord == (Coord{-1, -1}) || currCoord.Diff(prevCoord) == (Coord{-1, 0}) {
		sum += traverse(letters, sumCoords(currCoord, Coord{-1, 0}), currCoord, searchIdx + 1)
	}
	if prevCoord == (Coord{-1, -1}) || currCoord.Diff(prevCoord) == (Coord{1, 0}) {
		sum += traverse(letters, sumCoords(currCoord, Coord{1, 0}), currCoord, searchIdx + 1)
	}
	if prevCoord == (Coord{-1, -1}) || currCoord.Diff(prevCoord) == (Coord{-1, 1}) {
		sum += traverse(letters, sumCoords(currCoord, Coord{-1, 1}), currCoord, searchIdx + 1)
	}
	if prevCoord == (Coord{-1, -1}) || currCoord.Diff(prevCoord) == (Coord{0, 1}) {
		sum += traverse(letters, sumCoords(currCoord, Coord{0, 1}), currCoord, searchIdx + 1)
	}
	if prevCoord == (Coord{-1, -1}) || currCoord.Diff(prevCoord) == (Coord{1, 1}) {
		sum += traverse(letters, sumCoords(currCoord, Coord{1, 1}), currCoord, searchIdx + 1)
	}
	return sum
}