package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day2.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	part1 := 0
	part2 := 0

	for scanner.Scan() {
		levelsStr := strings.Fields(scanner.Text())
		levels := make([]int, len(levelsStr))
		for idx, s := range levelsStr {
			n, err := strconv.Atoi(s)

			if err != nil {
				log.Fatal(err)
			}
			levels[idx] = n
		}

		diff := 0
		// fmt.Println(levels)
		part1++
		for i := 1; i < len(levels); i++ {
			newDiff := levels[i] - levels[i-1]
			// fmt.Println(newDiff)
			if (diff > 0 && newDiff < 0) || (diff < 0 && newDiff > 0) || newDiff == 0 {
				part1--
				break
			}
			diff = newDiff
			newDiff = max(newDiff, -newDiff)
			if !(newDiff >= 1 && newDiff <= 3) {
				part1--
				break
			}
		}
		// fmt.Println(levels)
		isValid := validateLevels(levels)
		if isValid {
			part2++
		}
		// fmt.Printf("%v: %t\n", levels, isValid)
	}

	fmt.Printf("part1: %d\n", part1)
	fmt.Printf("part2: %d\n", part2)
}

func validateLevels(levels []int) bool {
	isValid, _ := validateLevelsRec(levels, true)
	return isValid
}

func validateLevelsRec(levels []int, canFail bool) (bool, int) {
	diff := 0
	fmt.Println(levels)

	for i := 0; i < len(levels) - 1; i++ {
		isValid, newDiff := validateDiff(levels[i], levels[i+1], diff)
		if (isValid) {
			diff = newDiff
			continue
		}

		if (canFail) {
			// fmt.Println("recursing")
			if i > 0 {
				left, _ := validateLevelsRec(slices.Delete(slices.Clone(levels), i-1, i), false)
				if left {
					return true, -1
				}
			}
			center, _ := validateLevelsRec(slices.Delete(slices.Clone(levels), i, i+1), false)
			if (center) {
				return center, -1
			}
			right, _ := validateLevelsRec(slices.Delete(slices.Clone(levels), i+1, i+2), false)

			return right, -1
		} else {
			return false, i
		}
		
	}

	return true, -1

}

func validateDiff(n1, n2, prevDiff int) (bool, int) {
	newDiff := n2 - n1
	if (prevDiff > 0 && newDiff < 0) || (prevDiff < 0 && newDiff > 0) || newDiff == 0 {
		// fmt.Printf("%d %d ||| %d vs %d\n",n1, n2, newDiff, prevDiff)
		return false, 0
	}
	newDiffAbs := max(newDiff, -newDiff)
	if (newDiffAbs < 1 || newDiffAbs > 3) {
		// fmt.Printf("%d %d ||| %d vs %d\n",n1, n2, newDiff, prevDiff)
		return false, 0
	}
	// fmt.Printf("%d %d ||| %d vs %d\n",n1, n2, newDiff, prevDiff)
	return true, newDiff
}

