package main

import (
	"day5/lines"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"time"
)

func main() {
	lines := lines.MustParse("data", "\n")
	var ids []int
	seatParser := makeSeatParser()
	for _, rule := range lines {
		ids = append(ids, seatParser(rule))
	}
	sortedIDs := sort.IntSlice(ids)
	sortedIDs.Sort()

	start := time.Now()
	fmt.Println(part1(sortedIDs))
	fmt.Println(part2(sortedIDs))
	t := time.Now()
	fmt.Println(t.Sub(start))

}

func part1(ids []int) int {
	return ids[len(ids)-1]
}

func part2(ids []int) int {
	expected := 0
	for ix := range ids {
		expected = ix + ids[0]
		if expected != ids[ix] {
			break
		}
	}
	return expected
}

func makeSeatParser() func(string) int {
	expr0 := regexp.MustCompile("[FL]")
	expr1 := regexp.MustCompile("[BR]")
	return func(rule string) int {
		rule = expr0.ReplaceAllString(rule, "0")
		rule = expr1.ReplaceAllString(rule, "1")
		id, err := strconv.ParseUint(rule, 2, 10)
		if err != nil {
			panic(err)
		}
		return int(id)
	}
}
