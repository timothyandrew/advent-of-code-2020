package impl

import "testing"

func TestCodeToSeat(t *testing.T) {
	if seat := codeToSeat("FBFBBFFRLR"); seat.row != 44 || seat.col != 5 {
		t.Errorf("Row/col incorrect: %d/%d", seat.row, seat.col)
	}
}
