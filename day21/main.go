package main

import (
	"day21/lines"
	"fmt"
	"sort"
	"strings"
	"time"
	"unicode"
)

func main() {
	lns := lines.MustParse("data", "\n")
	start := time.Now()
	foods, allIngredients := parseAllFoods(lns)
	fmt.Println(part1(foods, allIngredients))
	fmt.Println(part2(foods))
	stop := time.Now()
	fmt.Println(stop.Sub(start))
}

func part1(foods map[string]StringSlice, allIngredients StringSlice) int {
	for _, ingredients := range foods {
		allIngredients = allIngredients.removeIfPresent(ingredients[0])
	}
	return len(allIngredients)
}

func part2(foods map[string]StringSlice) string {
	var fds Foods
	for allergen, ingredients := range foods {
		fds = append(fds, Food{allergen: allergen, ingredient: ingredients[0]})
	}
	sort.Sort(fds)

	var result []string
	for _, food := range fds {
		result = append(result, food.ingredient)
	}
	return strings.Join(result, ",")
}

func parseAllFoods(lns []string) (map[string]StringSlice, StringSlice) {
	foods := make(map[string]StringSlice)
	var allIngredients StringSlice
	for _, line := range lns {
		ingredients, allergens := parseFood(line)
		allIngredients = append(allIngredients, ingredients...)
		for _, a := range allergens {
			_, ok := foods[a]
			if !ok {
				foods[a] = append(foods[a], ingredients...)
			} else {
				foods[a] = ingredients.commonValues(foods[a])
			}
		}
	}
	for {
		allFound := true
		for _, ingredients := range foods {
			if len(ingredients) != 1 {
				allFound = false
				break
			}
		}
		if allFound {
			break
		}
		for allergen, ingredients := range foods {
			if len(ingredients) == 1 {
				for a := range foods {
					if a == allergen {
						continue
					}
					foods[a] = foods[a].removeIfPresent(ingredients[0])
				}
			}
		}
	}
	return foods, allIngredients
}

func parseFood(food string) (StringSlice, StringSlice) {
	parts := strings.Split(food, " (contains ")
	ingredients := StringSlice(strings.Split(parts[0], " "))
	allergens := StringSlice(strings.Split(strings.TrimSuffix(parts[1], ")"), ", "))
	return ingredients, allergens
}

type StringSlice []string

func (s StringSlice) commonValues(s2 StringSlice) StringSlice {
	var same StringSlice
	for _, str := range s2 {
		if s.includes(str) {
			same = append(same, str)
		}
	}
	return same
}

func (s StringSlice) indexOf(str string) int {
	for ix, s := range s {
		if s == str {
			return ix
		}
	}
	return -1
}

func (s StringSlice) includes(str string) bool {
	return s.indexOf(str) != -1
}

func (s StringSlice) removeIfPresent(str string) StringSlice {
	for s.includes(str) {
		ix := s.indexOf(str)
		(s)[ix] = s[len(s)-1]
		s = s[0 : len(s)-1]
	}
	return s
}

type Food struct {
	allergen   string
	ingredient string
}
type Foods []Food

// copied straight from github
func (f Foods) Len() int      { return len(f) }
func (f Foods) Swap(i, j int) { f[i], f[j] = f[j], f[i] }
func (f Foods) Less(i, j int) bool {
	iRunes := []rune(f[i].allergen)
	jRunes := []rune(f[j].allergen)

	max := len(iRunes)
	if max > len(jRunes) {
		max = len(jRunes)
	}

	for idx := 0; idx < max; idx++ {
		ir := iRunes[idx]
		jr := jRunes[idx]

		lir := unicode.ToLower(ir)
		ljr := unicode.ToLower(jr)

		if lir != ljr {
			return lir < ljr
		}

		// the lowercase runes are the same, so compare the original
		if ir != jr {
			return ir < jr
		}
	}

	return false
}
