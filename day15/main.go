package main

import (
	"day15/lines"
	"fmt"
	"time"
)

func main() {
	numbers := lines.MustParseInts("data", ",")
	start := time.Now()
	fmt.Println(part(numbers, 2020))
	fmt.Println(part(numbers, 30000000))
	stop := time.Now()
	fmt.Println(stop.Sub(start))
}

func part(numbers []int, until int) int {
	numcntr := make(map[int][]int)
	for ix, n := range numbers {
		numcntr[n] = []int{ix + 1}
	}
	num := numbers[len(numbers)-1]
	for ix := len(numbers) + 1; ix <= until; ix++ {
		if len(numcntr[num]) <= 1 {
			num = 0
		} else {
			num = numcntr[num][0] - numcntr[num][1]
		}
		if len(numcntr[num]) > 0 {
			numcntr[num] = []int{ix, numcntr[num][0]}
		} else {
			numcntr[num] = []int{ix}
		}
	}
	return num
}
