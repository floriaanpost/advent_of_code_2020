package main

import (
	"day1/lines"
	"fmt"
	"time"
)

const year = 2020

func main() {
	numbers := lines.MustParseInts("data")

	start := time.Now()
	fmt.Println(part1(numbers))
	fmt.Println(part2(numbers))
	t := time.Now()
	fmt.Println(t.Sub(start))
}

func part1(numbers []int) int {
	for i, n1 := range numbers {
		for _, n2 := range numbers[i+1:] {
			if n1+n2 == year {
				return n1 * n2
			}
		}
	}
	panic("value not found")
}

func part2(numbers []int) int {
	for i, n1 := range numbers {
		for j, n2 := range numbers[i+1:] {
			for _, n3 := range numbers[j+1:] {
				if n1+n2+n3 == year {
					return n1 * n2 * n3
				}
			}
		}
	}
	panic("value not found")
}
