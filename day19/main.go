package main

import (
	"day19/lines"
	"fmt"
	"regexp"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(part1("data1"))
	fmt.Println(part2("data2"))
	stop := time.Now()
	fmt.Println(stop.Sub(start))
}

func part1(file string) int {
	blocks := lines.MustParse(file, "\n\n")
	ruleRegex := buildRuleRegex(blocks[0], 1)
	dataLines := strings.Split(blocks[1], "\n")
	count := 0
	for _, line := range dataLines {
		if ruleRegex.MatchString(line) {
			count++
		}
	}
	return count
}

func part2(file string) int {
	blocks := lines.MustParse(file, "\n\n")
	ruleRegex := buildRuleRegex(blocks[0], 5)
	dataLines := strings.Split(blocks[1], "\n")
	count := 0
	for _, line := range dataLines {
		if ruleRegex.MatchString(line) {
			count++
		}
	}
	return count
}

var ruleNumExpr = regexp.MustCompile(`^[0-9]+`)
var endRuleExpr = regexp.MustCompile(`[ab]`)

func findRule(rulenum string, rules string) string {
	rule := regexp.MustCompile("(\n|^)" + rulenum + ":.+(\n|$)")
	return strings.TrimPrefix(strings.TrimSuffix(rule.FindString(rules), "\n"), "\n")
}

func parseRule(rulenum string, rules string, seen map[string]int, maxRecursion int) string {
	r := findRule(rulenum, rules)
	ruleparts := strings.Split(r, ": ")
	rule := ruleparts[1]
	end := endRuleExpr.FindString(rule)
	if end != "" {
		return end
	}
	orRules := strings.Split(rule, " | ")
	regex := "("
	for ix, orRule := range orRules {
		ruleNums := strings.Split(orRule, " ")
		for _, num := range ruleNums {
			newSeen := make(map[string]int)
			for k, v := range seen {
				newSeen[k] = v
			}
			newSeen[num]++
			if newSeen[num] > maxRecursion {
				regex += "c" // does not match anything
			} else {
				regex += parseRule(num, rules, newSeen, maxRecursion)
			}
		}
		if ix != len(orRules)-1 {
			regex += "|"
		}
	}
	regex += ")"
	return regex
}

func buildRuleRegex(ruleBlock string, maxRecursion int) *regexp.Regexp {
	r := "^" + parseRule("0", ruleBlock, make(map[string]int), maxRecursion) + "$"
	regexrule := regexp.MustCompile(r)
	return regexrule
}
