package main

import (
	"day2/lines"
	"fmt"
	"strings"
	"time"
)

func main() {
	lns := lines.MustParse("data", "\n")

	start := time.Now()
	fmt.Println(part1(lns))
	fmt.Println(part2(lns))
	t := time.Now()
	fmt.Println(t.Sub(start))
}

func part1(lns []string) int {
	checkFunc := func(p pass, c chan bool) {
		count := strings.Count(p.password, string(p.c))
		c <- count >= p.min && count <= p.max
	}
	return mapLines(lns, checkFunc)
}

func part2(lns []string) int {
	checkFunc := func(p pass, c chan bool) {
		runes := []rune(p.password)
		count := 0
		if runes[p.min-1] == p.c {
			count++
		}
		if runes[p.max-1] == p.c {
			count++
		}
		c <- count == 1
	}
	return mapLines(lns, checkFunc)
}

func mapLines(lns []string, check func(pass, chan bool)) int {
	correct := 0
	c := make(chan bool, len(lns))
	for _, line := range lns {
		p := mustParsePass(line)
		check(p, c)
	}
	for rec := 0; rec < len(lns); rec++ {
		if <-c {
			correct++
		}
	}
	close(c)
	return correct
}

type pass struct {
	min      int
	max      int
	c        rune
	password string
}

func mustParsePass(line string) pass {
	var min, max int
	var c rune
	var password string
	_, err := fmt.Sscanf(line, "%d-%d %c: %s", &min, &max, &c, &password)
	if err != nil {
		panic(err)
	}
	return pass{min: min, max: max, c: c, password: password}
}
