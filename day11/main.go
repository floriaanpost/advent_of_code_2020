package main

import (
	"day11/lines"
	"fmt"
	"time"
)

func main() {
	lns := lines.MustParse("data", "\n")
	b := parseBoatLayout(lns)

	start := time.Now()

	fmt.Println(part1(b))
	fmt.Println(part2(b))

	t := time.Now()
	fmt.Println(t.Sub(start))
}

const (
	floor = iota
	empty
	occupied
	wall
)

type pos int
type boat [][]pos

func part1(b boat) int {
	for {
		newboat := b.copy()
		b.mapPositions(func(x, y int, p pos) bool {
			if p == floor {
				return true
			}
			occ := b.countOccupied(x, y, 1)
			if occ == 0 {
				newboat.setPos(x, y, occupied)
			}
			if occ >= 4 {
				newboat.setPos(x, y, empty)
			}
			return true
		})

		if newboat.equal(b) {
			break
		}
		b = newboat
	}
	return b.totalOccupied()
}

func part2(b boat) int {
	for {
		newboat := b.copy()
		b.mapPositions(func(x, y int, p pos) bool {
			if p == floor {
				return true
			}
			occ := b.countOccupied(x, y, -1)
			if occ == 0 {
				newboat.setPos(x, y, occupied)
			}
			if occ >= 5 {
				newboat.setPos(x, y, empty)
			}
			return true
		})

		if newboat.equal(b) {
			break
		}
		b = newboat
	}
	return b.totalOccupied()
}

func parseBoatLayout(lns []string) boat {
	var b boat
	seats := map[rune]pos{'.': floor, 'L': empty, '#': occupied}
	for _, line := range lns {
		var r []pos
		for _, v := range line {
			r = append(r, seats[v])
		}
		b = append(b, r)
	}
	return b
}

func (b boat) getPos(x, y int) pos {
	if x < 0 || y < 0 || y >= len(b) || x >= len(b[0]) {
		return wall
	}
	return b[y][x]
}

func (b boat) setPos(x, y int, p pos) {
	if p != empty && p != occupied {
		panic("can only set to emtpy or occupied")
	}
	b[y][x] = p
}

func (b boat) countOccupied(posx, posy int, maxDist int) int {
	count := 0
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}
			var p pos
			dist := 1
			for p == floor && (maxDist == -1 || dist <= maxDist) {
				p = b.getPos(posx+dist*x, posy+dist*y)
				if p == occupied {
					count++
				}
				dist++
			}
		}
	}
	return count
}

func (b boat) copy() boat {
	var newboat boat
	for _, row := range b {
		newrow := make([]pos, len(row))
		copy(newrow, row)
		newboat = append(newboat, newrow)
	}
	return newboat
}

func (b boat) equal(b2 boat) bool {
	return b.mapPositions(func(x, y int, p pos) bool {
		return b.getPos(x, y) == b2.getPos(x, y)
	})
}

func (b boat) mapPositions(mapFunc func(int, int, pos) bool) bool {
	for y, row := range b {
		for x, p := range row {
			if !mapFunc(x, y, p) {
				return false
			}
		}
	}
	return true
}

func (b boat) totalOccupied() int {
	count := 0
	b.mapPositions(func(_, _ int, p pos) bool {
		if p == occupied {
			count++
		}
		return true
	})
	return count
}
