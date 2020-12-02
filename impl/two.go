package impl

import (
	"fmt"
	"regexp"

	"timothyandrew.net/advent-2020/util"
)

func countRuneInString(c rune, s string) int {
	runes := []rune(s)

	count := 0
	for _, r := range runes {
		if r == c {
			count++
		}
	}

	return count
}

func Two() {
	re := regexp.MustCompile("(\\d+)\\-(\\d+)\\s(\\w):\\s(\\w+)")
	lines := util.ReadFileRegexMatches("input/2.txt", *re)

	// Part 1
	valid := 0
	for _, line := range lines {
		// Ugh
		lower, upper, char, s := util.ToInt(line[0]), util.ToInt(line[1]), line[2], line[3]
		count := countRuneInString([]rune(char)[0], s)

		if count >= lower && count <= upper {
			valid++
		}
	}

	fmt.Println("PART 1: Valid", valid)

	// Part 2
	valid = 0
	for _, line := range lines {
		index1, index2, char, s := util.ToInt(line[0]), util.ToInt(line[1]), []rune(line[2])[0], line[3]

		matches := 0

		if rune(s[index1-1]) == char {
			matches++
		}

		if rune(s[index2-1]) == char {
			matches++
		}

		if matches == 1 {
			valid++
		}
	}

	fmt.Println("PART 2: Valid", valid)
}
