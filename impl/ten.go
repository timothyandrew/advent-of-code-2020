package impl

import (
	"fmt"
	"sort"

	"timothyandrew.net/advent-2020/util"
)

func possibleNextSteps(n int, options []int) (result []int) {
	for i := 0; i < len(options); i++ {
		diff := options[i] - n
		if diff > 0 {
			if diff <= 3 {
				result = append(result, options[i])
			} else {
				break
			}
		}
	}

	return
}

func calcDifferences(s []int) map[int]int {
	result := make(map[int]int)

	for i := 0; i < len(s)-1; i++ {
		diff := s[i+1] - s[i]
		result[diff]++
	}

	return result
}

func buildArrangementsCountByIndex(steps []int) map[int]int {
	results := make(map[int]int)

	// Seed the lookup table with arrangement counts we know
	results[166] = 1
	results[163] = 1

	// Starting from the end of the list, compute the arrangement count of
	// a given element as the sum of the arrangement counts of each of the
	// possible next steps starting from that element. The two seed values above
	// ensure that we have a foundation to start from. Eventually we get to
	// the beginning of the list, and the arrangement count of 0 is our result.
	for i := len(steps) - 3; i >= 0; i-- {
		total := 0
		for _, nextStep := range possibleNextSteps(steps[i], steps[i+1:]) {
			total += results[nextStep]
		}
		results[steps[i]] = total
	}

	return results
}

func Ten() {
	lines := util.ReadFileInts("input/10.txt")
	sort.Ints(lines)
	lines = append([]int{0}, lines...)
	lines = append(lines, 166)

	fmt.Println("PART 1:", calcDifferences(lines))
	fmt.Println("PART 2:", buildArrangementsCountByIndex(lines)[0])
}
