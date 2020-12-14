package main

import (
	"day6/lines"
	"fmt"
	"strings"
	"time"
)

type question map[string]bool

func main() {
	groups := lines.MustParse("data", "\n\n")
	var questions [][]question
	for _, group := range groups {
		persons := strings.Split(group, "\n")
		var gq []question
		for _, p := range persons {
			q := make(question)
			for _, c := range p {
				q[string(c)] = true
			}
			gq = append(gq, q)
		}
		questions = append(questions, gq)
	}

	start := time.Now()
	fmt.Println(part1(questions))
	fmt.Println(part2(questions))
	t := time.Now()
	fmt.Println(t.Sub(start))
}

func part1(questions [][]question) int {
	var answeredYes []question
	for _, group := range questions {
		q := make(question)
		for _, person := range group {
			for c := range person {
				q[c] = true
			}
		}
		answeredYes = append(answeredYes, q)
	}
	return countYes(answeredYes)
}

func part2(questions [][]question) int {
	var answeredYes []question
	for _, group := range questions {
		q := make(question)
		for _, c := range "abcdefghijklmnopqrstuvwxyz" {
			pass := true
			for _, person := range group {
				if !person[string(c)] {
					pass = false
					break
				}
			}
			q[string(c)] = pass
		}
		answeredYes = append(answeredYes, q)
	}
	return countYes(answeredYes)
}

func countYes(yesQuestions []question) int {
	count := 0
	for _, q := range yesQuestions {
		for _, pass := range q {
			if pass {
				count++
			}
		}
	}
	return count
}
