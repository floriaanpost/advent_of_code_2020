package main

import (
	"day7/lines"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

type childBags map[string]int

func main() {
	lns := lines.MustParse("data", "\n")
	parse := mustMakeBagParser()
	bags := make(map[string]childBags)
	for _, l := range lns {
		name, cbs := parse(l)
		bags[name] = cbs
	}

	start := time.Now()
	fmt.Println(part1(bags))
	fmt.Println(part2(bags))
	t := time.Now()
	fmt.Println(t.Sub(start))
}

func part1(bags map[string]childBags) int {
	bagTypes := []string{"shiny gold"}
	var foundParentTypes []string
	for len(bagTypes) > 0 {
		var nextBagTypes []string
		for parent, children := range bags {
			for _, bt := range bagTypes {
				if children[bt] > 0 {
					if isNew(foundParentTypes, parent) {
						foundParentTypes = append(foundParentTypes, parent)
						nextBagTypes = append(nextBagTypes, parent)
					}
				}
			}
		}
		bagTypes = nextBagTypes
	}
	return len(foundParentTypes)
}

func part2(bags map[string]childBags) int {
	return countChildren(bags, "shiny gold") - 1 // don't count gold bag, so -1
}

func mustMakeBagParser() func(string) (string, childBags) {
	bagTypeRegex := regexp.MustCompile(`^.+ contain`)
	childBagsRegex := regexp.MustCompile(`([0-9]+.+?(,|\.))`)
	childBagCountRegex := regexp.MustCompile(`^[0-9]+`)
	childBagNameRegex := regexp.MustCompile(`([a-z]+ )+[a-z]+$`)
	bagSuffixRegex := regexp.MustCompile(` bags*(,|\.| contain)*`)

	return func(line string) (string, childBags) {
		bagType := bagTypeRegex.FindString(line)
		bagType = bagSuffixRegex.ReplaceAllString(bagType, "")

		matches := childBagsRegex.FindAllString(line, -1)
		cbs := make(childBags)
		for _, m := range matches {
			m := bagSuffixRegex.ReplaceAllString(m, "")
			count := childBagCountRegex.FindString(m)
			name := childBagNameRegex.FindString(m)
			c, err := strconv.ParseInt(count, 10, 0)
			if err != nil {
				panic(err)
			}
			cbs[name] = int(c)
		}
		return bagType, cbs
	}
}

func isNew(found []string, cur string) bool {
	for _, v := range found {
		if v == cur {
			return false
		}
	}
	return true
}

func countChildren(bags map[string]childBags, name string) int {
	count := 1 // count the bag itself
	for child, amount := range bags[name] {
		count += amount * countChildren(bags, child)
	}
	return count
}
