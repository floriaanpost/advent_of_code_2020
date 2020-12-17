package main

import (
	"day14/lines"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	lns := lines.MustParse("data", "\n")
	start := time.Now()
	fmt.Println(part1(lns))
	fmt.Println(part2(lns))
	stop := time.Now()
	fmt.Println(stop.Sub(start))
}

func part1(lns []string) uint64 {
	mem := make(map[uint64]uint64)
	var bitSet, bitReset uint64
	for _, l := range lns {
		if strings.HasPrefix(l, "mask") {
			mask := strings.TrimPrefix(l, "mask = ")
			bitSet, bitReset = parseMask(mask)
			continue
		}
		address, value := parseInstruction(l)
		mem[address] = (value & bitReset) | bitSet
	}
	return sumMem(mem)
}

func part2(lns []string) uint64 {
	mem := make(map[uint64]uint64)
	var masks []uint64
	var xbits uint64
	for _, l := range lns {
		if strings.HasPrefix(l, "mask") {
			mask := strings.TrimPrefix(l, "mask = ")
			xbits, masks = allMasksUint(mask)
			continue
		}
		address, value := parseInstruction(l)
		for _, m := range masks {
			mem[(address|m)&(^((m & xbits) ^ xbits))] = value
		}
	}
	return sumMem(mem)
}

func allMasksUint(mask string) (uint64, []uint64) {
	xstr := strings.ReplaceAll(mask, "1", "0")
	xstr = strings.ReplaceAll(xstr, "X", "1")
	xbits, _ := strconv.ParseUint(xstr, 2, 36)
	masks := allMasks(mask)
	var result []uint64
	for _, m := range masks {
		bitSet, _ := parseMask(m)
		result = append(result, bitSet)
	}
	return xbits, result
}

func allMasks(mask string) []string {
	if !strings.Contains(mask, "X") {
		return []string{mask}
	}
	mask1 := strings.Replace(mask, "X", "1", 1)
	mask2 := strings.Replace(mask, "X", "0", 1)
	var masks []string
	masks = append(masks, allMasks(mask1)...)
	masks = append(masks, allMasks(mask2)...)
	return masks
}

func parseMask(mask string) (uint64, uint64) {
	mreset := strings.ReplaceAll(mask, "X", "1")
	bitReset, _ := strconv.ParseUint(mreset, 2, 36)
	mset := strings.ReplaceAll(mask, "X", "0")
	bitSet, _ := strconv.ParseUint(mset, 2, 36)
	return bitSet, bitReset
}

func parseInstruction(instruction string) (uint64, uint64) {
	var address, value uint64
	fmt.Sscanf(instruction, "mem[%d] = %d", &address, &value)
	return address, value
}

func sumMem(mem map[uint64]uint64) uint64 {
	var sum uint64
	for _, val := range mem {
		sum += val
	}
	return sum
}
