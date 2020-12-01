package impl

import (
	"fmt"

	"timothyandrew.net/advent-2020/util"
)

func arrayWithout(arr []int, without int) []int {
	var result []int
	for _, i := range arr {
		if i != without {
			result = append(result, i)
		}
	}
	return result
}

func solve2(target int, ints []int) (x, y int, success bool) {
	lookup := make(map[int]bool)
	for _, i := range ints {
		lookup[i] = true
	}

	for _, i := range ints {
		rem := target - i

		if lookup[rem] {
			return i, rem, true
		}
	}

	return 0, 0, false
}

func One() {
	ints := util.ReadFileInts("input/1.txt")

	// Part 1
	x, y, succ := solve2(2020, ints)
	if succ {
		fmt.Println("PART 1: Found a solution!", x, y, x*y)
	}

	// Part 2
	for _, i := range ints {
		rem := 2020 - i

		x, y, succ = solve2(rem, arrayWithout(ints, i))
		if succ {
			fmt.Println("PART 2: Found a solution!", i, x, y, i*x*y)
			return
		}
	}
}
