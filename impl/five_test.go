package impl

import "testing"

func TestCodeToSeat(t *testing.T) {
	if row, col := codeToSeat("FBFBBFFRLR"); row != 44 || col != 5 {
		t.Errorf("Row/col incorrect: %d/%d", row, col)
	}
}
