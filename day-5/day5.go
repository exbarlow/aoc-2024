package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)
type Rule struct {
	left, right string
}

func (r Rule) Applies (m map[string]int) bool {
	return m[r.left] > 0 && m[r.right] > 0
}

func (r Rule) Allows (m map[string]int) bool {
	return m[r.left] < m[r.right]
}

func (r Rule) ToKey () string {
	return fmt.Sprintf("%s|%s", r.left, r.right)
}

func main() {
	b, err := os.ReadFile("day5.txt")

	if err != nil {
		log.Fatal(err)
	}

	inputStr := string(b)

	strs := strings.Split(inputStr,"\n\n")

	rulesStrs := strings.Split(strs[0], "\n")
	pageStrs := strings.Split(strs[1], "\n")

	var rules []Rule
	for _, rulesStr := range rulesStrs {
		ruleParts := strings.Split(rulesStr, "|")
		rules = append(rules, Rule{ruleParts[0], ruleParts[1]})
	}

	var pages [][]string
	for _, pageStr := range pageStrs {
		pageNums := strings.Split(pageStr, ",")
		pages = append(pages, pageNums)
	}

	ruleDict := make(map[string]Rule)

	for _, rule := range rules {
		ruleDict[rule.ToKey()] = rule
	}

	
	var unorderedPages [][]string
	part1 := 0
	for _, page := range pages {
		if isOrdered(page, rules) {
			part1 += getMiddlePageNum(page)
		} else {
			unorderedPages = append(unorderedPages, page)
		}
	}
	part2 := 0
	for _, page := range unorderedPages {
		sortSliceByMap(page, ruleDict)
		part2 += getMiddlePageNum(page)
	}

	fmt.Printf("part1: %d\n", part1)
	fmt.Printf("part2: %d\n", part2)
}

func getMiddlePageNum(page []string) int {
	l := len(page)
	// fmt.Printf("l, page: %d, %v\n", l, page)
	i, err := strconv.Atoi(page[l/2])
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func sortSliceByMap(page []string, ruleDict map[string]Rule) {
	slices.SortFunc(page, func(a, b string) int {
		key := Rule{a, b}.ToKey()
		if ruleDict[key] != (Rule{"",""}) {
			return -1
		}
		key = Rule{b, a}.ToKey()
		if ruleDict[key] != (Rule{"",""}) {
			return 1
		}
		return 0
	})
}

func isOrdered(page []string, rules []Rule) bool {
	pagePosMap := make(map[string]int)

	for idx, p := range page {
		pagePosMap[p] = idx + 1
	}

	for _, rule := range rules {
		if rule.Applies(pagePosMap) && !rule.Allows(pagePosMap) {
			return false
		}
	}
	return true
}