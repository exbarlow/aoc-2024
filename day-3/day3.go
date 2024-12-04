package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	input, err := os.ReadFile("day3.txt")

	if err != nil {
		log.Fatal(err)
	}

	r, err := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)
	r2, err2 := regexp.Compile(`(?:^|do\(\)|don't\(\))`)

	if err != nil {
		log.Fatal(err)
	}

	if err2 != nil {
		log.Fatal(err2)
	}

	inputStr := string(input)
	s := r2.Split(inputStr, -1)
	s2 := r2.FindAllString(inputStr, -1)

	if len(s) != len(s2) {
		log.Fatalf("expected lens to match, given %d and %d", len(s), len(s2))
	}

	part2 := 0
	for i := range s {
		if s2[i] == "don't()" {
			continue
		}
		matches := r.FindAllStringSubmatch(s[i], -1)
		for _, matchPair := range matches {
			if len(matchPair) != 3 {
				log.Fatalf("expected 3 matches, got %d (%v)\n", len(matchPair), matchPair)
			}
			val1, err1 := strconv.Atoi(matchPair[1])
			val2, err2 := strconv.Atoi(matchPair[2])
	
			if err1 != nil || err2 != nil {
				log.Fatal("did not expect err")
			}
			part2 += val1 * val2
		}
	}

	matches := r.FindAllSubmatch(input, -1)
	part1 := 0
	for _, matchPair := range matches {
		if len(matchPair) != 3 {
			log.Fatalf("expected 3 matches, got %d (%v)\n", len(matchPair), matchPair)
		}
		val1, err1 := strconv.Atoi(string(matchPair[1]))
		val2, err2 := strconv.Atoi(string(matchPair[2]))

		if err1 != nil || err2 != nil {
			log.Fatal("did not expect err")
		}
		part1 += val1 * val2

	}
	fmt.Printf("part1: %d\n", part1)
	fmt.Printf("part1: %d\n", part2)
}