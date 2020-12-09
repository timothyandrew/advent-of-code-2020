package impl

import (
	"fmt"

	"timothyandrew.net/advent-2020/util"
)

func findSumPair(nums []int, target int) (int, int, bool) {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return nums[i], nums[j], true
			}
		}
	}

	return 0, 0, false
}

func Nine() {
	var input []int = util.ReadFileInts("input/9.txt")
	var result int

	for i := 25; i < len(input); i++ {
		_, _, ok := findSumPair(input[i-25:i], input[i])

		if !ok {
			fmt.Println("PART 1:", input[i])
			result = input[i]
		}
	}

	for i := 0; i < len(input); i++ {
		var acc []int

		for j := i; j < len(input); j++ {
			acc = append(acc, input[j])

			if util.SumSlice(acc) == result {
				fmt.Println("PART 2:", util.MinSlice(acc), util.MaxSlice(acc))
				return
			}

			if util.SumSlice(acc) > result {
				break
			}
		}
	}
}
