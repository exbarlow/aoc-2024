package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	ADD = "+"
	MULT = "*"
	CONCAT = "||"
)

type Equation struct {
	target int
	numbers []int
}

func main() {
	fileName := "day7.txt"
	fmt.Printf("part1: %d\n", part1(fileName))
	fmt.Printf("part2: %d\n", part2(fileName))
}

func part2(f string) int {
	b, err := os.ReadFile(f); if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")
	var equations []Equation
	for _, line := range lines {
		equations = append(equations, buildEquation(line))
	}

	ans := 0
	for _, eq := range equations {
		if eq.OperateWithConcat(0, 0, ADD) || eq.OperateWithConcat(0, 0, MULT) || eq.OperateWithConcat(0, 0, CONCAT) {
			ans += eq.target
		}
	}
	return ans
}

func part1(f string) int {
	b, err := os.ReadFile(f); if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")
	var equations []Equation
	for _, line := range lines {
		equations = append(equations, buildEquation(line))
	}

	ans := 0
	for _, eq := range equations {
		if eq.Operate(0, 0, ADD) || eq.Operate(0, 0, MULT) {
			ans += eq.target
		}
	}
	return ans
}

func (e* Equation) OperateWithConcat(currVal, rIdx int, op string) bool {
	if rIdx >= len(e.numbers) || rIdx < 0 {
		log.Fatalf("Invalid index (%d) for equation with numbers %v\n", rIdx, e.numbers)
	}
	var nextVal int
	if op == ADD {
		nextVal = currVal + e.numbers[rIdx]
	} else if op == MULT {
		nextVal = currVal * e.numbers[rIdx]
	} else if op == CONCAT {
		if currVal == 0 {
			nextVal = e.numbers[rIdx]
		} else {
			var err error
			nextVal, err = strconv.Atoi(strconv.Itoa(currVal) + strconv.Itoa(e.numbers[rIdx])); if err != nil {
				log.Fatal(err)
			}
		}
	} else {
		log.Fatalf("Expected valid operation, got %s\n", op);
	}

	if rIdx == len(e.numbers) - 1 {
		return nextVal == e.target
	}

	return e.OperateWithConcat(nextVal, rIdx + 1, ADD) || e.OperateWithConcat(nextVal, rIdx + 1, MULT) || e.OperateWithConcat(nextVal, rIdx + 1, CONCAT)
}

func (e *Equation) Operate(currVal, rIdx int, op string) bool {
	if rIdx >= len(e.numbers) || rIdx < 0 {
		log.Fatalf("Invalid index (%d) for equation with numbers %v\n", rIdx, e.numbers)
	}
	var nextVal int
	if op == ADD {
		nextVal = currVal + e.numbers[rIdx]
	} else if op == MULT {
		nextVal = currVal * e.numbers[rIdx]
	} else {
		log.Fatalf("Expected valid operation, got %s\n", op);
	}

	if rIdx == len(e.numbers) - 1 {
		return nextVal == e.target
	}

	return e.Operate(nextVal, rIdx + 1, ADD) || e.Operate(nextVal, rIdx + 1, MULT)
}

// TODO: refactor to use error handling
func buildEquation(s string) Equation {
	parts := strings.Split(s, ":")
	if len(parts) != 2 {
		log.Fatalf("Split operation created more parts than expected. Original string: (%s), parts: (%v)", s, parts)
	}
	target, err := strconv.Atoi(parts[0]); if err != nil {
		log.Fatal(err)
	}
	var numbers []int
	for _, str := range strings.Fields(parts[1]) {
		n, err := strconv.Atoi(str); if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, n)
	}
	return Equation{target, numbers}
}


