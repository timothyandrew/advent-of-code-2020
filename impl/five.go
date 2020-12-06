package impl

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strings"

	"timothyandrew.net/advent-2020/util"
)

func codeToSeat(code string) seat {
	rowCode := strings.Split(code[:7], "")
	colCode := strings.Split(code[7:], "")

	lower, upper := float64(0), float64(127)
	for _, row := range rowCode {
		diff := math.Ceil((upper - lower) / 2)
		if row == "F" {
			upper -= diff
		} else if row == "B" {
			lower += diff
		}
	}

	row := int(lower)

	lower, upper = float64(0), float64(7)
	for _, row := range colCode {
		diff := math.Ceil((upper - lower) / 2)
		if row == "L" {
			upper -= diff
		} else if row == "R" {
			lower += diff
		}
	}

	col := int(lower)

	return seat{row: row, col: col}
}

type seat struct {
	col int
	row int
}

func Five() {
	re := regexp.MustCompile("^([FB]{7}[LR]{3})$")
	lines := util.ReadFileLinesRegexMatches("input/5.txt", *re)

	var seatIDs []int
	seats := make(map[seat]seat)
	for _, line := range lines {
		seat := codeToSeat(line[0])
		seatIDs = append(seatIDs, (seat.row*8)+seat.col)
		seats[seat] = seat
	}

	sort.Ints(seatIDs)

	fmt.Println("PART 1:", seatIDs[len(seatIDs)-1])

	fmt.Println("\nPART 2")
	for row := 0; row < 128; row++ {
		for col := 0; col < 8; col++ {
			if _, ok := seats[seat{row: row, col: col}]; !ok {
				fmt.Printf("Seat with row %v and col %v is missing!\n", row, col)
			}
		}
	}
}
