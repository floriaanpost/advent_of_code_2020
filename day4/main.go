package main

import (
	"day4/lines"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	lns := lines.MustParse("data", "\n\n")

	start := time.Now()
	fmt.Println(part1(lns))
	fmt.Println(part2(lns))
	t := time.Now()
	fmt.Println(t.Sub(start))
}

func part1(lns []string) int {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	checkPassport := func(p passport, c chan bool) {
		pass := true
		for _, field := range requiredFields {
			v := p.getField(field)
			if v == nil {
				pass = false
				break
			}
		}
		c <- pass
	}
	return validCount(lns, checkPassport)
}

func part2(lns []string) int {
	fieldValidators := map[string]func(string) bool{
		"byr": between(1920, 2002),
		"iyr": between(2010, 2020),
		"eyr": between(2020, 2030),
		"hgt": height(),
		"hcl": regex("^#[0-9a-f]{6}$"),
		"ecl": regex("^(amb|blu|brn|gry|grn|hzl|oth)$"),
		"pid": regex("^[0-9]{9}$"),
	}

	checkPassport := func(p passport, c chan bool) {
		for field, validator := range fieldValidators {
			v := p.getField(field)
			if v == nil {
				c <- false
				return
			}
			valid := validator(*v)
			if !valid {
				c <- false
				return
			}
		}
		c <- true
	}
	return validCount(lns, checkPassport)
}

func validCount(lns []string, checkPassport func(passport, chan bool)) int {
	count := 0
	c := make(chan bool, len(lns))
	for _, line := range lns {
		checkPassport(passport(line), c)
	}
	for rec := 0; rec < len(lns); rec++ {
		if <-c {
			count++
		}
	}
	return count
}

type passport string

func (p passport) getField(field string) *string {
	re, err := regexp.Compile(field + ":[A-z0-9#]+")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	match := re.FindString(string(p))
	if match == "" {
		return nil
	}
	parts := strings.Split(match, ":")
	return &parts[1]
}

func between(lbound, ubound int) func(string) bool {
	return func(s string) bool {
		i, err := strconv.ParseInt(s, 10, 0)
		if err != nil {
			return false
		}
		return i >= int64(lbound) && i <= int64(ubound)
	}
}

func regex(restr string) func(string) bool {
	re := regexp.MustCompile(restr)
	return func(s string) bool {
		return re.MatchString(s)
	}
}

func height() func(string) bool {
	validRe := regex("^[0-9]+(cm|in)$")
	numRe := regexp.MustCompile("[0-9]+")
	validCM := between(150, 193)
	validIN := between(59, 76)
	return func(s string) bool {
		if !(validRe(s)) {
			return false
		}
		cm := strings.Contains(s, "cm")
		n := numRe.FindString(s)
		if cm {
			return validCM(n)
		}
		return validIN(n)
	}
}
