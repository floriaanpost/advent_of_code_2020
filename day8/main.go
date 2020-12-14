package main

import (
	"day8/lines"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	instructions := lines.MustParse("data", "\n")

	start := time.Now()
	fmt.Println(part1(instructions))
	fmt.Println(part2(instructions))
	t := time.Now()
	fmt.Println(t.Sub(start))
}

func part1(instructions []string) int {
	acc, _ := mustExec(instructions, 0, 0, make(map[int]bool))
	return int(acc)
}

func part2(instructions []string) int {
	var result int
	for i, oldLine := range instructions {
		nop := strings.Contains(oldLine, "nop")
		jmp := strings.Contains(oldLine, "jmp")
		if !nop && !jmp {
			continue
		}
		var newLine string
		if nop {
			newLine = strings.Replace(oldLine, "nop", "jmp", 1)
		} else {
			newLine = strings.Replace(oldLine, "jmp", "nop", 1)
		}
		instructions[i] = newLine
		if acc, exit := mustExec(instructions, 0, 0, make(map[int]bool)); exit {
			result = acc
			break
		}
		instructions[i] = oldLine
	}
	return result
}

func mustExec(instructions []string, linenr int, acc int, accessedLines map[int]bool) (int, bool) {
	accessedLines[linenr] = true
	l := instructions[linenr]
	p := strings.Split(l, " ")
	cmd := p[0]
	val, err := strconv.ParseInt(p[1], 10, 0)
	if err != nil {
		panic(err)
	}
	switch cmd {
	case "acc":
		linenr++
		acc += int(val)
	case "jmp":
		linenr += int(val)
	case "nop":
		linenr++
	}
	if accessedLines[linenr] {
		return acc, false
	}
	if int(linenr) >= len(instructions) {
		return acc, true
	}
	return mustExec(instructions, linenr, acc, accessedLines)
}
