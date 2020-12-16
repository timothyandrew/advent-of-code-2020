package impl

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/stat/combin"
	"timothyandrew.net/advent-2020/util"
)

type Mask struct {
	maskOne  int64
	maskZero int64
}

func (m Mask) apply(val int64) int64 {
	val = val & m.maskZero
	val = val | m.maskOne
	return val
}

type SeaportCPUState struct {
	mask Mask
	mem  map[int64]int64
}

func (cpu *SeaportCPUState) updateMask(s string) {
	maskOne, err := strconv.ParseInt(strings.ReplaceAll(s, "X", "0"), 2, 64)
	if err != nil {
		panic("Invalid binary string")
	}

	maskZero, err := strconv.ParseInt(strings.ReplaceAll(s, "X", "1"), 2, 64)
	if err != nil {
		panic("Invalid binary string")
	}

	cpu.mask = Mask{maskOne: maskOne, maskZero: maskZero}
}

func (cpu *SeaportCPUState) updateMem(addr, val int64) {
	cpu.mem[addr] = cpu.mask.apply(val)
}

func (cpu SeaportCPUState) SumMemValues() (sum int64) {
	for _, val := range cpu.mem {
		sum += val
	}
	return
}

type SeaportCPUV2State struct {
	SeaportCPUState
	memMasks []Mask
}

func (cpu *SeaportCPUV2State) updateMask(s string) {
	re := regexp.MustCompile("X")

	var xIndexes []int
	var lens []int

	for _, match := range re.FindAllIndex([]byte(s), -1) {
		xIndexes = append(xIndexes, match[0])
		lens = append(lens, 2)
	}

	var masks []Mask
	for _, combo := range combin.Cartesian(lens) {
		maskOneStr := []rune(s)
		maskZeroStr := []rune("111111111111111111111111111111111111")

		// Set up maskOne and maskZero so that ((addr & maskOne) | maskZero)
		// applies the incoming mask, and forces the correct `X` values.
		for i, index := range xIndexes {
			val := []rune(fmt.Sprint(combo[i]))[0]
			maskOneStr[index] = val

			if val == '0' {
				maskZeroStr[index] = '0'
			}
		}

		masks = append(masks, Mask{
			maskOne:  util.FromBinaryString(string(maskOneStr)),
			maskZero: util.FromBinaryString(string(maskZeroStr)),
		})
	}

	cpu.memMasks = masks
}

func (cpu *SeaportCPUV2State) updateMem(addr, val int64) {
	for _, mask := range cpu.memMasks {
		cpu.mem[mask.apply(addr)] = val
	}
}

func Fourteen() {
	lines := util.FileToLines("input/14.txt")
	cpu := SeaportCPUState{mask: Mask{maskOne: 0, maskZero: 1}, mem: make(map[int64]int64)}
	cpuV2 := SeaportCPUV2State{SeaportCPUState: SeaportCPUState{mask: Mask{maskOne: 0, maskZero: 1}, mem: make(map[int64]int64)}}

	for _, line := range lines {
		maskRe := regexp.MustCompile("mask = (\\S+)")
		memRe := regexp.MustCompile("mem\\[(\\d+)\\] = (\\d+)")

		if matches := maskRe.FindStringSubmatch(line); matches != nil {
			cpu.updateMask(matches[1])
			cpuV2.updateMask(matches[1])
		}

		if matches := memRe.FindStringSubmatch(line); matches != nil {
			cpu.updateMem(util.ToInt64(matches[1]), util.ToInt64(matches[2]))
			cpuV2.updateMem(util.ToInt64(matches[1]), util.ToInt64(matches[2]))
		}
	}

	fmt.Println("Part 1:", cpu.SumMemValues())
	fmt.Println("Part 2:", cpuV2.SumMemValues())
}
