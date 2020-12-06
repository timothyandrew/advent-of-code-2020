package impl

import (
	"fmt"
	"strings"

	"timothyandrew.net/advent-2020/util"
)

func Six() {
	s := util.FileToString("input/6.txt")

	lines := strings.Split(s, "\n\n")
	counts := 0

	for _, groupStr := range lines {
		group := strings.Join(strings.Split(groupStr, "\n"), "")

		answer := make(map[string]bool)

		for _, member := range strings.Split(group, "") {
			answer[member] = true
		}

		counts += len(answer)
	}

	fmt.Println("PART 1", counts)

	counts = 0
	for _, groupStr := range lines {
		groups := strings.Split(groupStr, "\n")
		groupCount := len(groups)

		answer := make(map[string]int)
		for _, group := range groups {
			for _, member := range strings.Split(group, "") {
				answer[member]++
			}
		}

		for _, c := range answer {
			if c == groupCount {
				counts++
			}
		}
	}

	fmt.Println("PART 2", counts)
}
