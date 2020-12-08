package impl

import (
	"fmt"
	"strings"

	"timothyandrew.net/advent-2020/util"
)

type Instruction struct {
	opcode string
	n      int
	index  int
}

type CPUState struct {
	acc          int
	ip           int
	instructions []Instruction
	seen         map[Instruction]bool
	stopped      bool
	success      bool
}

func parse(s string, index int) Instruction {
	matches := strings.Split(s, " ")
	opcode, nStr := matches[0], matches[1]
	n := util.ToInt(nStr)

	return Instruction{opcode: opcode, n: n, index: index}
}

func tick(state CPUState) CPUState {
	current := state.instructions[state.ip]

	if state.ip == len(state.instructions) {
		state.stopped = true
		state.success = true
		fmt.Println("Attempting to run an instruction after the last instruction. Stopping…\n")
		return state
	}

	if _, ok := state.seen[current]; ok {
		state.stopped = true
		fmt.Printf("Seeing instruction (%v) for the second time! Stopping…\n", current)
		return state
	}

	// I'm told this copies the struct; does it?
	next := state
	next.seen[current] = true

	switch current.opcode {
	case "acc":
		next.acc += current.n
		next.ip++
	case "jmp":
		next.ip += current.n
	case "nop":
		next.ip++
	default:
		panic("Invalid opcode")
	}

	return next
}

func Eight() {
	lines := util.FileToLines("input/8.txt")

	var instructions []Instruction
	for index, line := range lines {
		instructions = append(instructions, parse(line, index))
	}

	state := CPUState{instructions: instructions, seen: make(map[Instruction]bool)}

	for !state.stopped {
		state = tick(state)
	}

	fmt.Println("PART 1: ", state.acc)
}
