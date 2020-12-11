package impl

import (
	"fmt"
	"strings"

	"timothyandrew.net/advent-2020/util"
)

type cell struct {
	state string
}

func newCell(i rune) (c cell) {
	switch i {
	case '#':
		c = cell{state: "occupied"}
	case 'L':
		c = cell{state: "empty"}
	case '.':
		c = cell{state: "floor"}
	default:
		panic("invalid cell")
	}

	return c
}

func (c cell) isOccupied() bool {
	return c.state == "occupied"
}

func (c cell) isEmpty() bool {
	return c.state == "empty"
}

func (c cell) isSeat() bool {
	return c.state != "floor"
}

type board struct {
	cells [][]cell
}

func newBoard(lines []string) board {
	var cells [][]cell
	for _, line := range lines {
		var current []cell

		for _, c := range strings.Split(line, "") {
			current = append(current, newCell([]rune(c)[0]))
		}

		cells = append(cells, current)
	}

	return board{cells}
}

func (b board) adjacentOccupied(x, y int) (occupied int) {
	if y > 0 {
		// Left
		if b.cells[x][y-1].isOccupied() {
			occupied++
		}
	}

	if x > 0 {
		// Top
		if b.cells[x-1][y].isOccupied() {
			occupied++
		}
	}

	if x > 0 && y > 0 {
		// Top-left
		if b.cells[x-1][y-1].isOccupied() {
			occupied++
		}
	}

	if x < len(b.cells)-1 {
		// Bottom
		if b.cells[x+1][y].isOccupied() {
			occupied++
		}
	}

	if y < len(b.cells[0])-1 {
		// Right
		if b.cells[x][y+1].isOccupied() {
			occupied++
		}
	}

	if x < len(b.cells)-1 && y < len(b.cells[0])-1 {
		// Bottom Right
		if b.cells[x+1][y+1].isOccupied() {
			occupied++
		}
	}

	if x > 0 && y < len(b.cells[0])-1 {
		// Top right
		if b.cells[x-1][y+1].isOccupied() {
			occupied++
		}
	}

	if x < len(b.cells)-1 && y > 0 {
		// Bottom left
		if b.cells[x+1][y-1].isOccupied() {
			occupied++
		}
	}

	return
}

func (b board) tick(part2 bool) board {
	var cells [][]cell

	for i, row := range b.cells {
		var current []cell

		for j, c := range row {
			occ := b.adjacentOccupied(i, j)
			emptyFactor := 4

			if part2 {
				occ = b.occupiedVisible(i, j)
				emptyFactor = 5
			}

			if c.isEmpty() && occ == 0 {
				current = append(current, cell{"occupied"})
			} else if c.isOccupied() && occ >= emptyFactor {
				current = append(current, cell{"empty"})
			} else {
				current = append(current, c)
			}
		}

		cells = append(cells, current)
	}

	return board{cells}
}

func (b board) isEqual(other board) bool {
	for i := 0; i < len(b.cells); i++ {
		for j := 0; j < len(b.cells[0]); j++ {
			if b.cells[i][j] != other.cells[i][j] {
				return false
			}
		}
	}

	return true
}

func (b board) forEachCell(f func(cell)) {
	for i := 0; i < len(b.cells); i++ {
		for j := 0; j < len(b.cells[0]); j++ {
			f(b.cells[i][j])
		}
	}
}

func (b board) occupiedVisible(x, y int) (occupied int) {
	// Top
	if x > 0 {
		for i := x - 1; i >= 0; i-- {
			if b.cells[i][y].isOccupied() {
				occupied++
			}

			if b.cells[i][y].isSeat() {
				break
			}
		}
	}

	// Left
	if y > 0 {
		for i := y - 1; i >= 0; i-- {
			if b.cells[x][i].isOccupied() {
				occupied++
			}

			if b.cells[x][i].isSeat() {
				break
			}
		}
	}

	// Bottom
	if x < len(b.cells)-1 {
		for i := x + 1; i < len(b.cells); i++ {
			if b.cells[i][y].isOccupied() {
				occupied++
			}

			if b.cells[i][y].isSeat() {
				break
			}
		}
	}

	// Right
	if y < len(b.cells[0])-1 {
		for i := y + 1; i < len(b.cells[0]); i++ {
			if b.cells[x][i].isOccupied() {
				occupied++
			}

			if b.cells[x][i].isSeat() {
				break
			}
		}
	}

	// Top-Left
	if y > 0 && x > 0 {
		for i, j := x-1, y-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
			if b.cells[i][j].isOccupied() {
				occupied++
			}

			if b.cells[i][j].isSeat() {
				break
			}
		}
	}

	// Top-Right
	if y < len(b.cells[0])-1 && x > 0 {
		for i, j := x-1, y+1; i >= 0 && j < len(b.cells[0]); i, j = i-1, j+1 {
			if b.cells[i][j].isOccupied() {
				occupied++
			}

			if b.cells[i][j].isSeat() {
				break
			}
		}
	}

	// Bottom-Right
	if y < len(b.cells[0])-1 && x < len(b.cells)-1 {
		for i, j := x+1, y+1; i < len(b.cells) && j < len(b.cells[0]); i, j = i+1, j+1 {
			if b.cells[i][j].isOccupied() {
				occupied++
			}

			if b.cells[i][j].isSeat() {
				break
			}
		}
	}

	// Bottom-Left
	if y > 0 && x < len(b.cells)-1 {
		for i, j := x+1, y-1; i < len(b.cells) && j >= 0; i, j = i+1, j-1 {
			if b.cells[i][j].isOccupied() {
				occupied++
			}

			if b.cells[i][j].isSeat() {
				break
			}
		}
	}

	return
}

func Eleven() {
	lines := util.FileToLines("input/11.txt")
	board := newBoard(lines)

	last := board
	next := last.tick(false)

	for !next.isEqual(last) {
		last = next
		next = last.tick(false)
	}

	occupied := 0
	next.forEachCell(func(c cell) {
		if c.isOccupied() {
			occupied++
		}
	})

	fmt.Println("PART 1:", occupied)

	board = newBoard(lines)

	last = board
	next = last.tick(true)

	for !next.isEqual(last) {
		last = next
		next = last.tick(true)
	}

	occupied = 0
	next.forEachCell(func(c cell) {
		if c.isOccupied() {
			occupied++
		}
	})

	fmt.Println("PART 2:", occupied)
}
