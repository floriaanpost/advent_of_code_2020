package main

import (
	"day18/lines"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	lns := lines.MustParse("data", "\n")
	fmt.Println(part1(lns))
	fmt.Println(part2(lns))
}

func part1(lns []string) int64 {
	var sum int64
	for _, l := range lns {
		result, _ := strconv.ParseInt(calculate(l, "order"), 0, 0)
		sum += result
	}
	return sum
}

func part2(lns []string) int64 {
	var sum int64
	for _, l := range lns {
		result, _ := strconv.ParseInt(calculate(l, "+"), 0, 0)
		sum += result
	}
	return sum
}

var bracketExpr = regexp.MustCompile(`\([0-9]+ (\+|\*) [0-9]+( (\+|\*) [0-9]+)*\)`)
var firstExpr = regexp.MustCompile(`^[0-9]+ (\+|\*) [0-9]+`)
var operators = regexp.MustCompile(` (\+|\*) `)
var addExpr = regexp.MustCompile(`[0-9]+ \+ [0-9]+`)
var multiExpr = regexp.MustCompile(`[0-9]+ \* [0-9]+`)

func calculate(math string, order string) string {
	expr := bracketExpr.FindString(math)
	if expr != "" {
		inner := strings.TrimPrefix(expr, "(")
		inner = strings.TrimSuffix(inner, ")")
		solution := calculate(inner, order)
		math = strings.Replace(math, expr, solution, 1)
		if bracketExpr.MatchString(math) {
			math = calculate(math, order)
			if !operators.MatchString(math) {
				return math
			}
		}
	}
	switch order {
	case "order":
		expr = firstExpr.FindString(math)
	case "+":
		expr = addExpr.FindString(math)
		if expr == "" {
			expr = multiExpr.FindString(math)
		}
	}
	nums := operators.Split(expr, 2)
	operator := strings.TrimSpace(operators.FindString(expr))
	n1, _ := strconv.ParseInt(nums[0], 0, 0)
	n2, _ := strconv.ParseInt(nums[1], 0, 0)
	var n int64
	switch operator {
	case "+":
		n = n1 + n2
	case "*":
		n = n1 * n2
	}
	math = strings.Replace(math, expr, fmt.Sprint(n), 1)
	if operators.MatchString(math) {
		return calculate(math, order)
	}
	return math
}
