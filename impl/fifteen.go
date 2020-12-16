package impl

import "fmt"

type recitationState struct {
	spoken       []int // Numbers spoken, in order
	spokenAtTurn map[int][]int
}

func (r recitationState) lastSpoken() int {
	return r.spoken[len(r.spoken)-1]
}

func (r recitationState) currentTurn() int { return len(r.spoken) }

func (r *recitationState) speak(n int) {
	r.spoken = append(r.spoken, n)

	if spokenAt, ok := r.spokenAtTurn[n]; ok && len(spokenAt) == 2 {
		r.spokenAtTurn[n] = []int{r.currentTurn(), spokenAt[0]}
	} else if ok && len(spokenAt) == 1 {
		r.spokenAtTurn[n] = []int{r.currentTurn(), spokenAt[0]}
	} else {
		r.spokenAtTurn[n] = []int{r.currentTurn()}
	}
}

func (r *recitationState) takeTurn() {
	if spokenAt, ok := r.spokenAtTurn[r.lastSpoken()]; ok && len(spokenAt) == 2 {
		r.speak(spokenAt[0] - spokenAt[1])
	} else {
		r.speak(0)
	}
}

func Fifteen() {
	state := recitationState{
		spoken:       []int{2, 0, 1, 9, 5, 19},
		spokenAtTurn: map[int][]int{2: {1}, 0: {2}, 1: {3}, 9: {4}, 5: {5}, 19: {6}},
	}

	for len(state.spoken) < 2020 {
		state.takeTurn()
	}

	fmt.Println("PART 1:", state.lastSpoken())

	state = recitationState{
		spoken:       []int{2, 0, 1, 9, 5, 19},
		spokenAtTurn: map[int][]int{2: {1}, 0: {2}, 1: {3}, 9: {4}, 5: {5}, 19: {6}},
	}

	for len(state.spoken) < 30000000 {
		state.takeTurn()
	}

	fmt.Println("PART 2:", state.lastSpoken())
}
