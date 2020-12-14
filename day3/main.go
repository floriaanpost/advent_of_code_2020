package main

import (
	"day3/lines"
	"fmt"
	"time"
)

type pos bool

const (
	empty = false
	tree  = true
)

type slope [][]pos

func main() {
	s := fileToSlope("data")

	start := time.Now()
	fmt.Println(part1(s))
	fmt.Println(part2(s))
	t := time.Now()
	fmt.Println(t.Sub(start))
}

func part1(s slope) int {
	return s.countTrees(3, 1)
}

func part2(s slope) int {
	c1 := s.countTrees(1, 1)
	c2 := s.countTrees(3, 1)
	c3 := s.countTrees(5, 1)
	c4 := s.countTrees(7, 1)
	c5 := s.countTrees(1, 2)
	return c1 * c2 * c3 * c4 * c5
}

func fileToSlope(filename string) slope {
	lns := lines.MustParse(filename, "\n")
	var s slope
	for _, line := range lns {
		var row []pos
		for _, c := range line {
			if c == '#' {
				row = append(row, tree)
			} else {
				row = append(row, empty)
			}
		}
		s = append(s, row)
	}
	return s
}

func (s slope) readPos(posx, posy int) pos {
	posx %= len(s[0])
	return s[posy][posx]
}

func (s slope) countTrees(dx, dy int) int {
	posx := 0
	count := 0
	for posy := 0; posy < len(s); posy += dy {
		if s.readPos(posx, posy) == tree {
			count++
		}
		posx += dx
	}
	return count
}
