package main

import (
	"day13/lines"
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

func part1(lns []string) int64 {
	now, _ := strconv.ParseInt(lns[0], 0, 0)
	busses, _ := mustParseBusses(lns[1])
	var bus int64
	earliest := now
	for bus == 0 {
		for _, b := range busses {
			if earliest%b == 0 {
				twait := earliest - now
				return b * twait
			}
		}
		earliest++
	}
	panic("bus not found!")
}

func part2(lns []string) int64 {
	busses, offsets := mustParseBusses(lns[1])
	var t int64 = busses[0]
	var dt int64 = busses[0]
	for i := 1; i < len(busses); i++ {
		bus := busses[i]
		offset := offsets[i]
		for (t+offset)%bus != 0 {
			t += dt
		}
		dt = dt * busses[i]
	}
	return t
}

func mustParseBusses(line string) ([]int64, []int64) {
	busdata := strings.Split(line, ",")
	var busses []int64
	var offsets []int64
	for ix, v := range busdata {
		if v == "x" {
			continue
		}
		busnr, _ := strconv.ParseInt(v, 0, 0)
		busses = append(busses, busnr)
		offsets = append(offsets, int64(ix))
	}
	return busses, offsets
}
