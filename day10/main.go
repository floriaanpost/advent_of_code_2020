package main

import (
	"day10/lines"
	"fmt"
	"sort"
	"time"
)

func main() {
	values := lines.MustParseInts("data")
	adapters := sort.IntSlice(values)
	adapters.Sort()
	adapters = append([]int{0}, adapters...)                 // add wall yoltage before
	adapters = append(adapters, adapters[len(adapters)-1]+3) // add device yoltage after

	start := time.Now()
	fmt.Println(part1(adapters))
	fmt.Println(part2(adapters))
	t := time.Now()
	fmt.Println(t.Sub(start))
}

func part1(adapters []int) int {
	var diffs []int
	for ix := range adapters[1:] {
		diffs = append(diffs, adapters[ix+1]-adapters[ix])
	}
	count := make(map[int]int)
	for _, d := range diffs {
		count[d]++
	}
	return count[1] * count[3]
}

func part2(adapters []int) uint64 {
	poss := make(map[int]uint64)
	poss[0] = 1
	for ix := 1; ix < len(adapters); ix++ {
		ixs := prevPossibleAdaptersIxs(adapters, ix)
		var p uint64
		for _, j := range ixs {
			p += poss[j]
		}
		poss[ix] = p
	}
	return poss[len(adapters)-1]
}

func prevPossibleAdaptersIxs(adapters []int, ix int) []int {
	c := adapters[ix]
	var diff int
	var prev []int
	ix--
	for diff <= 3 {
		if ix <= 0 {
			prev = append(prev, ix)
			break
		}
		prev = append(prev, ix)
		ix--
		diff = c - adapters[ix]
	}
	return prev
}
