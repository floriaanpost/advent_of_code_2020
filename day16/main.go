package main

import (
	"day16/lines"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	parts := lines.MustParse("data", "\n\n")
	ranges := parseRanges(parts[0])
	myTicket := parseTickets(parts[1])[0]
	otherTickets := parseTickets(parts[2])
	start := time.Now()
	fmt.Println(part1(ranges, otherTickets))
	fmt.Println(part2(ranges, otherTickets, myTicket))
	stop := time.Now()
	fmt.Println(stop.Sub(start))
}

func part1(ranges []numRange, tickets [][]int) int {
	var invalid []int
	for _, ticket := range tickets {
		_, nums := isValid(ranges, ticket)
		invalid = append(invalid, nums...)
	}
	count := 0
	for _, n := range invalid {
		count += n
	}
	return count
}

type opt struct {
	name string
	ixs  []int
}
type opts []opt

func (o opts) Len() int           { return len(o) }
func (o opts) Swap(i, j int)      { o[i], o[j] = o[j], o[i] }
func (o opts) Less(i, j int) bool { return len(o[i].ixs) < len(o[j].ixs) }

func part2(ranges []numRange, tickets [][]int, myTicket []int) int {
	var validTickets [][]int
	for _, ticket := range tickets {
		if ok, _ := isValid(ranges, ticket); ok {
			validTickets = append(validTickets, ticket)
		}
	}
	var options opts
	numcnt := len(tickets[0])
	for _, r := range ranges {
		var o opt
		for ix := 0; ix < numcnt; ix++ {
			valid := true
			for _, ticket := range validTickets {
				if !r.withinRange(ticket[ix]) {
					valid = false
					continue
				}
			}
			if valid {
				o.name = r.name
				o.ixs = append(o.ixs, ix)
				continue
			}
		}
		options = append(options, o)
	}
	sort.Sort(options)
	var seen []int
	result := make(map[string]int)
	for _, o := range options {
		for _, i := range o.ixs {
			hasSeen := false
			for _, ix := range seen {
				if i == ix {
					hasSeen = true
				}
			}
			if !hasSeen {
				result[o.name] = i
				seen = append(seen, i)
			}
		}
	}
	multi := 1
	for field, ix := range result {
		if strings.Contains(field, "departure") {
			multi *= myTicket[ix]
		}
	}
	return multi
}

type numRange struct {
	name   string
	ranges []struct {
		min int
		max int
	}
}

func (n *numRange) withinRange(num int) bool {
	for _, r := range n.ranges {
		if num >= r.min && num <= r.max {
			return true
		}
	}
	return false
}

func parseRanges(part string) []numRange {
	lines := strings.Split(part, "\n")
	var ranges []numRange
	for _, line := range lines {
		var ran numRange
		pts := strings.Split(line, ": ")
		rngs := strings.Split(pts[1], " or ")
		ran.name = pts[0]
		for _, r := range rngs {
			nums := strings.Split(r, "-")
			n1, _ := strconv.ParseInt(nums[0], 0, 0)
			n2, _ := strconv.ParseInt(nums[1], 0, 0)
			ran.ranges = append(ran.ranges, struct {
				min int
				max int
			}{int(n1), int(n2)})
		}
		ranges = append(ranges, ran)
	}
	return ranges
}

func parseTickets(part string) [][]int {
	lines := strings.Split(part, "\n")
	var tickets [][]int
	for _, line := range lines[1:] {
		var nums []int
		strs := strings.Split(line, ",")
		for _, s := range strs {
			n, _ := strconv.ParseInt(s, 0, 0)
			nums = append(nums, int(n))
		}
		tickets = append(tickets, nums)
	}
	return tickets
}

func isValid(ranges []numRange, ticket []int) (bool, []int) {
	var invalid []int
	for _, num := range ticket {
		valid := false
		for _, r := range ranges {
			if r.withinRange(num) {
				valid = true
				break
			}
		}
		if !valid {
			invalid = append(invalid, num)
		}
	}
	return len(invalid) == 0, invalid
}
