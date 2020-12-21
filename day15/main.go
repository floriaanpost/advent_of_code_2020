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
	lastseen := make([]int, until)
	for ix, n := range numbers {
		lastseen[n] = ix + 1
	}
	num := numbers[len(numbers)-1]
	for ix := len(numbers) + 1; ix <= until; ix++ {
		prev := lastseen[num]
		lastseen[num] = ix - 1
		if prev == 0 {
			num = 0
			continue
		}
		num = ix - 1 - prev
	}
	return num
}
