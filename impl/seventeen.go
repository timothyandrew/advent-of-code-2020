package impl

import (
	"fmt"
	"strings"

	"timothyandrew.net/advent-2020/util"
)

type coord3d struct {
	x, y, z int
}

type cube struct {
	active bool
}

type grid3d struct {
	cubes map[coord3d]cube
}

func (g grid3d) copy() grid3d {
	cubes := make(map[coord3d]cube)
	for k, v := range g.cubes {
		cubes[coord3d{k.x, k.y, k.z}] = cube{v.active}
	}
	return grid3d{cubes}
}

func (g grid3d) activeNeighbors(c coord3d) (count int) {
	for _, x := range []int{-1, 0, 1} {
		for _, y := range []int{-1, 0, 1} {
			for _, z := range []int{-1, 0, 1} {
				if x == 0 && y == 0 && z == 0 {
					continue
				}

				if cube, ok := g.cubes[coord3d{c.x + x, c.y + y, c.z + z}]; ok && cube.active {
					count++
				}
			}
		}
	}

	return
}

func (g *grid3d) expand() {
	min, max := 0, 0

	for coord := range g.cubes {
		if coord.x > max {
			max = coord.x
		}

		if coord.y > max {
			max = coord.y
		}

		if coord.z > max {
			max = coord.z
		}

		if coord.x < min {
			min = coord.x
		}

		if coord.y < min {
			min = coord.y
		}

		if coord.z < min {
			min = coord.z
		}
	}

	buffer := 3

	for x := min - buffer; x <= max+buffer; x++ {
		for y := min - buffer; y <= max+buffer; y++ {
			for z := min - buffer; z <= max+buffer; z++ {
				if _, ok := g.cubes[coord3d{x, y, z}]; !ok {
					g.cubes[coord3d{x, y, z}] = cube{false}
				}
			}
		}
	}
}

func (g grid3d) tick() grid3d {
	g.expand()
	result := g.copy()

	for coord, c := range g.cubes {
		neighbors := g.activeNeighbors(coord)

		if c.active && (neighbors < 2 || neighbors > 3) {
			result.cubes[coord] = cube{false}
		}

		if !c.active && (neighbors == 3) {
			result.cubes[coord] = cube{true}
		}
	}

	return result
}

func (g grid3d) countActive() (count int) {
	for _, c := range g.cubes {
		if c.active {
			count++
		}
	}
	return
}

func Seventeen() {
	lines := util.FileToLines("input/17.txt")
	cubes := make(map[coord3d]cube)
	for i, line := range lines {
		for j, c := range strings.Split(line, "") {
			cubes[coord3d{i, j, 0}] = cube{c == "#"}
		}
	}

	grid := grid3d{cubes}

	for i := 0; i < 6; i++ {
		grid = grid.tick()
	}

	fmt.Println("PART 1:", grid.countActive())

}
