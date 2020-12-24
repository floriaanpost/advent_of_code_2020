package main

import (
	"day22/lines"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	parts := lines.MustParse("data", "\n\n")
	start := time.Now()

	deck1 := parseDeck(parts[0])
	deck2 := parseDeck(parts[1])

	fmt.Println(part1(deck1, deck2))
	fmt.Println(part2(deck1, deck2))

	stop := time.Now()
	fmt.Println(stop.Sub(start))
}

func part1(deck1, deck2 Deck) int {
	_, winner := play(deck1, deck2, false)
	return winner.score()
}

func part2(deck1, deck2 Deck) int {
	_, winner := play(deck1, deck2, true)
	return winner.score()
}

var subGameNr int

func play(deck1, deck2 Deck, recursion bool) (bool, Deck) {
	var seenDecks1 []Deck
	for len(deck1) > 0 && len(deck2) > 0 {
		if hasSeenDeck(seenDecks1, deck1) {
			return true, deck1
		}
		seenDecks1 = append(seenDecks1, deck1)
		v1, v2 := deck1.drawTop(), deck2.drawTop()
		var player1Won bool
		if recursion && v1 <= len(deck1) && v2 <= len(deck2) {
			player1Won, _ = play(deck1[0:v1], deck2[0:v2], recursion)
		} else {
			player1Won = v1 > v2
		}
		if player1Won {
			deck1 = append(deck1, v1, v2)
		} else {
			deck2 = append(deck2, v2, v1)
		}
	}
	if len(deck1) == 0 {
		return false, deck2
	}
	return true, deck1
}

func parseDeck(input string) Deck {
	lns := strings.Split(input, "\n")
	var deck Deck
	for _, val := range lns[1:] {
		i, _ := strconv.ParseInt(val, 0, 0)
		deck = append(deck, int(i))
	}
	return deck
}

var counter int

func hasSeenDeck(decks []Deck, deck Deck) bool {
	for ix := range decks {
		if decks[ix].equal(deck) {
			return true
		}
	}
	return false
}

type Deck []int

func (d *Deck) drawTop() int {
	if len(*d) == 0 {
		return -1
	}
	first := (*d)[0]
	*d = (*d)[1:]
	return first
}

func (d *Deck) score() int {
	score := 0
	for ix, v := range *d {
		score += v * (len(*d) - ix)
	}
	return score
}

func (d *Deck) equal(d2 Deck) bool {
	if len(*d) != len(d2) {
		return false
	}
	for ix := range *d {
		if (*d)[ix] != d2[ix] {
			return false
		}
	}
	return true
}
