package main

import (
	"day17/lines"
	"fmt"
	"time"
)

// not very fast, some improvements could easily be made...
func main() {
	lns := lines.MustParse("data", "\n")
	start := time.Now()
	fmt.Println(simulate(lns, 3, 6))
	fmt.Println(simulate(lns, 4, 6))
	stop := time.Now()
	fmt.Println(stop.Sub(start))
}

func simulate(lns []string, dims, rounds int) int {
	s := readSpace(lns, dims)
	for i := 0; i < rounds; i++ {
		scopy := s.copy()
		s.forEachPos(func(coord []int) {
			alive := s.isAlive(coord)
			count := s.neigbourCount(coord)
			if alive && (count < 2 || count > 3) {
				scopy.setDead(coord)
			}
			if !alive && count == 3 {
				scopy.setAlive(coord)
			}
		})
		s = scopy
	}
	return len(s)
}

func readSpace(lns []string, dims int) space {
	var s space
	var extraDims []int
	for ix := 0; ix < dims-2; ix++ {
		extraDims = append(extraDims, 0)
	}
	for y, l := range lns {
		for x, c := range l {
			if c == '#' {
				coord := append([]int{x, y}, extraDims...)
				s.setAlive(coord)
			}
		}
	}
	return s
}

type space [][]int

func allCoords(min, max []int, prefix []int) [][]int {
	var coords [][]int
	for i := min[0] - 1; i <= max[0]+1; i++ {
		newPrefix := make([]int, len(prefix)+1)
		copy(newPrefix, append(prefix, i))
		if len(min) > 1 {
			coords = append(coords, allCoords(min[1:], max[1:], newPrefix)...)
		} else {
			coords = append(coords, newPrefix)
		}
	}
	return coords
}

func allNeighbours(coord []int, prefix []int) [][]int {
	var coords [][]int
	for offset := -1; offset <= 1; offset++ {
		newPrefix := make([]int, len(prefix)+1)
		copy(newPrefix, append(prefix, coord[0]+offset))
		if len(coord) > 1 {
			coords = append(coords, allNeighbours(coord[1:], newPrefix)...)
		} else {
			coords = append(coords, newPrefix)
		}
	}
	return coords
}

func equal(sl1, sl2 []int) bool {
	for ix := range sl1 {
		if sl1[ix] != sl2[ix] {
			return false
		}
	}
	return true
}

func (s *space) forEachPos(mapFunc func([]int)) {
	min, max := s.getBounds()
	all := allCoords(min, max, []int{})
	for _, coord := range all {
		mapFunc(coord)
	}
}

func (s *space) neigbourCount(coord []int) int {
	neighbours := allNeighbours(coord, []int{})
	count := 0
	for _, n := range neighbours {
		if !equal(coord, n) && s.isAlive(n) {
			count++
		}
	}
	return count
}

func (s *space) findIndex(coord []int) int {
	for ix, p := range *s {
		if equal(coord, p) {
			return ix
		}
	}
	return -1
}

func (s *space) isAlive(coord []int) bool {
	return s.findIndex(coord) != -1
}

func (s *space) setAlive(coord []int) {
	if s.isAlive(coord) {
		return
	}
	*s = append(*s, coord)
}

func (s *space) setDead(coord []int) {
	ix := s.findIndex(coord)
	if ix == -1 {
		return
	}
	(*s)[ix] = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
}

func (s *space) getBounds() ([]int, []int) {
	min := make([]int, len((*s)[0]))
	max := make([]int, len((*s)[0]))
	for _, p := range *s {
		for ix, c := range p {
			if c < min[ix] {
				min[ix] = c
			}
			if c > max[ix] {
				max[ix] = c
			}
		}
	}
	return min, max
}

func (s *space) copy() space {
	var newspace space
	for _, p := range *s {
		newp := make([]int, len(p))
		copy(newp, p)
		newspace = append(newspace, newp)
	}
	return newspace
}
