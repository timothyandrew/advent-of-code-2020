package impl

import (
	"reflect"
	"testing"
)

func TestSeaportCpu(t *testing.T) {
	cpu := SeaportCPUState{maskOne: 0, maskZero: 1, mem: make(map[int64]int64)}

	cpu.updateMask("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X")
	cpu.updateMem(8, 11)
	cpu.updateMem(7, 101)
	cpu.updateMem(8, 0)

	if eq := reflect.DeepEqual(cpu.mem, map[int64]int64{7: 101, 8: 64}); !eq {
		t.Error("Test failed!", cpu.mem)
	}
}
