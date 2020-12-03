package impl

import (
	"fmt"
	"regexp"
	"strings"

	"timothyandrew.net/advent-2020/util"
)

type coord struct {
	x int
	y int
}

type grid struct {
	width  int
	height int
}

func takeStep(c coord, g grid, step coord) coord {
	return coord{
		x: (c.x + step.x) % g.width,
		y: (c.y + step.y) % g.height,
	}
}

func traverse(lines [][]string) {
	trees := grid{len(lines[0]), len(lines)}

	slopes := []coord{{3, 1}, {1, 1}, {5, 1}, {7, 1}, {1, 2}}

	for _, slope := range slopes {
		step := coord{0, 0}

		obstacles := 0
		for step.y < (trees.height - 1) {
			step = takeStep(step, trees, slope)

			if lines[step.y][step.x] == "#" {
				obstacles++
			}
		}

		fmt.Printf("Slope %v: %v trees!\n", slope, obstacles)
	}
}

func Three() {
	re := regexp.MustCompile("([.#]+)")
	lines := util.ReadFileRegexMatches("input/3.txt", *re)

	for i, line := range lines {
		lines[i] = strings.Split(line[0], "")
	}

	traverse(lines)
}
