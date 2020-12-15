package main

import (
	"day12/lines"
	"fmt"
	"math"
	"time"
)

func main() {
	lns := lines.MustParse("data", "\n")
	cmds := linesToCmds(lns)

	start := time.Now()

	fmt.Println(part1(cmds))
	fmt.Println(part2(cmds))

	t := time.Now()
	fmt.Println(t.Sub(start))
}

func part1(cmds []*cmd) int {
	s := newShip(0, 0, 90)
	for _, c := range cmds {
		s.mustMove(c)
	}
	return int(math.Abs(float64(s.x)) + math.Abs(float64(s.y)))
}

func part2(cmds []*cmd) int {
	w := newWaypoint(10, 1)
	s := struct {
		y int
		x int
	}{0, 0}
	for _, c := range cmds {
		if c.instruction == 'F' {
			s.x += c.amount * w.x
			s.y += c.amount * w.y
		} else {
			w.mustMove(c)
		}
	}
	return int(math.Abs(float64(s.x)) + math.Abs(float64(s.y)))
}

type cmd struct {
	instruction rune
	amount      int
}

func linesToCmds(lns []string) []*cmd {
	var c rune
	var n int
	var cmds []*cmd
	for _, l := range lns {
		fmt.Sscanf(l, "%c%d", &c, &n)
		cmds = append(cmds, &cmd{amount: n, instruction: c})
	}
	return cmds
}

type ship struct {
	y    int
	x    int
	head int
}

func newShip(y, x, head int) *ship {
	return &ship{x: x, y: y, head: head}
}

func (s *ship) mustMove(c *cmd) {
	i := c.instruction
	n := c.amount
	if i == 'F' {
		switch s.head {
		case 0:
			i = 'N'
		case 90:
			i = 'E'
		case 180:
			i = 'S'
		case 270:
			i = 'W'
		default:
			panic(fmt.Sprintf("Unknown heading: %d", s.head))
		}
	}

	switch i {
	case 'N':
		s.x += n
	case 'E':
		s.y += n
	case 'S':
		s.x -= n
	case 'W':
		s.y -= n
	case 'L':
		s.head = (s.head + 360 - n) % 360
	case 'R':
		s.head = (s.head + n) % 360
	default:
		panic(fmt.Sprintf("Unknown instruction: %c", i))
	}
}

type waypoint struct {
	y int
	x int
}

func newWaypoint(y, x int) *waypoint {
	return &waypoint{x: x, y: y}
}

func (w *waypoint) mustMove(c *cmd) {
	i := c.instruction
	n := c.amount
	switch i {
	case 'N':
		w.x += n
	case 'E':
		w.y += n
	case 'S':
		w.x -= n
	case 'W':
		w.y -= n
	case 'L':
		w.rotate(-n)
	case 'R':
		w.rotate(n)
	default:
		panic(fmt.Sprintf("Unknown instruction: %c", i))
	}
}

func (w *waypoint) rotate(deg int) {
	beta := (float64(deg) / 360) * 2 * math.Pi
	w.x, w.y = int(math.Round(math.Cos(beta)*float64(w.x)-math.Sin(beta)*float64(w.y))), int(math.Round(math.Sin(beta)*float64(w.x)+math.Cos(beta)*float64(w.y)))
}
