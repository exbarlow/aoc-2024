package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("day1.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r, _ := regexp.Compile(`^(\d+) *(\d+)$`)

	scanner := bufio.NewScanner(file)

	leftNums := []int{}
	rightNums := make(map[int]int)

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
		// fmt.Printf("%d %d\n", leftNumParsed, rightNumParsed)

		leftNums = append(leftNums, leftNumParsed)
		rightNums[rightNumParsed] = rightNums[rightNumParsed] + 1
	}

	ans := 0

	for i := 0; i < len(leftNums); i++ {
		ans += leftNums[i] * rightNums[leftNums[i]]
	}

	fmt.Printf("Ans: %d\n", ans)
}