package main

import (
	"day9/lines"
	"fmt"
	"time"
)

const preambleSize = 25

func main() {
	numbers := lines.MustParseInts("data")
	start := time.Now()
	fmt.Println(part1(numbers))
	fmt.Println(part2(numbers))
	t := time.Now()
	fmt.Println(t.Sub(start))
}

func part1(numbers []int) int {
	correct := makeCodeChecker(numbers, preambleSize)
	var ix int
	for ix = preambleSize; ix < len(numbers); ix++ {
		if !correct(ix) {
			break
		}
	}
	return numbers[ix]
}

func part2(numbers []int) int {
	sumFunc := makeSummer(numbers, preambleSize, part1(numbers))
	var sumNumbers []int
	for ix := range numbers {
		if res, ok := sumFunc(ix); ok {
			sumNumbers = res
		}
	}
	min, max := int(^uint(0)>>1), 0 // 0 and largest value in int type
	for _, n := range sumNumbers {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min + max
}

func makeCodeChecker(code []int, preambleSize int) func(int) bool {
	return func(ix int) bool {
		if ix < preambleSize {
			panic("ix should be smaller than the preamble size")
		}
		preamble := code[ix-preambleSize : ix]
		for i, num1 := range preamble {
			for _, num2 := range preamble[i:] {
				if num1+num2 == code[ix] {
					return true
				}
			}
		}
		return false
	}
}

func makeSummer(code []int, preambleSize int, value int) func(int) ([]int, bool) {
	return func(ix int) ([]int, bool) {
		sum := 0
		var values []int
		for sum < value {
			val := code[ix]
			values = append(values, val)
			sum += val
			ix++
		}
		return values, sum == value && len(values) >= 2
	}
}
