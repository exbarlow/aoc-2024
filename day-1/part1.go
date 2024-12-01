package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func p1() {
	file, err := os.Open("day1.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r, _ := regexp.Compile(`^(\d+) *(\d+)$`)

	scanner := bufio.NewScanner(file)

	leftNums := []int{}
	rightNums := []int{}

	for scanner.Scan() {
		text := scanner.Text()
		strs := r.FindStringSubmatch(text)
		if len(strs) != 3 {
			log.Fatalf("Expected 3 matches, given %d", len(strs))
		}
		leftNumParsed, err := strconv.Atoi(strs[1])
		if (err != nil) {
			log.Fatal(err)
		}
		rightNumParsed, err := strconv.Atoi(strs[2])
		if (err != nil) {
			log.Fatal(err)
		}

		leftNums = append(leftNums, leftNumParsed)
		rightNums = append(rightNums, rightNumParsed)
	}

	slices.Sort(leftNums)
	slices.Sort(rightNums)

	if len(leftNums) != len(rightNums) {
		log.Fatalf("len(left): %d does not match len(right): %d", len(leftNums), len(rightNums))
	}

	ans := 0

	for i := 0; i < len(leftNums); i++ {
		diff := leftNums[i] - rightNums[i]
		ans += max(diff, -diff)
	}

	fmt.Printf("Ans: %d\n", ans)
}