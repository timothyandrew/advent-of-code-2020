package impl

import (
	"fmt"
	"regexp"
	"strings"

	"timothyandrew.net/advent-2020/util"
)

type Bag struct {
	color string
}

type Rule struct {
	bag      Bag
	contains map[Bag]int
}

func parseRule(rule string) Rule {
	re := regexp.MustCompile("(.*) bags contain (.* bags?,?\\s?)+")
	matches := re.FindStringSubmatch(rule)

	from, to := matches[1], matches[2]

	contains := make(map[Bag]int)
	for _, bagStr := range strings.Split(to, ",") {
		bagStr = strings.Trim(bagStr, "\n \t\r")

		if strings.HasPrefix(bagStr, "no") {

		} else {
			re = regexp.MustCompile("(\\d+) (.*) bags?")
			matches = re.FindStringSubmatch(bagStr)
			count, bag := util.ToInt(matches[1]), Bag{color: matches[2]}
			contains[bag] = count
		}
	}

	return Rule{bag: Bag{color: from}, contains: contains}
}

func countBagsContainedIn(bag Bag, lookup map[Bag]Rule) int {
	resolved := 1

	rule, ok := lookup[bag]
	if !ok {
		return resolved
	}

	for bag, count := range rule.contains {
		resolved += count * countBagsContainedIn(bag, lookup)
	}

	return resolved
}

func findRoots(bag Bag, lookup map[Bag][]Bag) map[Bag]bool {
	resolved := make(map[Bag]bool)
	resolved[bag] = true

	bags, ok := lookup[bag]
	if !ok {
		return resolved
	}

	for _, bag := range bags {
		for target := range findRoots(bag, lookup) {
			resolved[target] = true
		}
	}

	return resolved
}

func Seven() {
	lines := util.FileToLines("input/7.txt")

	containedBy := make(map[Bag][]Bag)

	for _, line := range lines {
		rule := parseRule(line)

		for c := range rule.contains {
			containedBy[c] = append(containedBy[c], rule.bag)
		}
	}

	x := findRoots(Bag{color: "shiny gold"}, containedBy)

	// Subtract 1 because a shiny gold bag can't contain itself
	fmt.Println("PART 1:", len(x)-1)

	contains := make(map[Bag]Rule)

	for _, line := range lines {
		rule := parseRule(line)
		contains[rule.bag] = rule
	}

	y := countBagsContainedIn(Bag{color: "shiny gold"}, contains)

	// Subtract 1 because a shiny gold bag can't contain itself
	fmt.Println("PART 2:", y-1)
}
