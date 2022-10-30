package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var matching = map[byte]byte{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

var syntaxPoints = map[byte]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

// Checks the syntax of a line and returns its score.
func syntaxCheck(line string) int {
	s := stack{}
	for _, c := range []byte(line) {
		switch c {
		case '(': // Opening chunck
			fallthrough
		case '[':
			fallthrough
		case '{':
			fallthrough
		case '<':
			s.push(c)
		case ')': // Closing chuck
			fallthrough
		case ']':
			fallthrough
		case '}':
			fallthrough
		case '>':
			opening := s.pop()
			if c != matching[opening] {
				return syntaxPoints[c]
			}
		}
	}
	return 0
}

var completionPoints = map[byte]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

// Autocompletes a line and returns its score.
func autocomplete(line string) int {
	s := stack{}
	for _, c := range []byte(line) {
		switch c {
		case '(': // Opening chunck
			fallthrough
		case '[':
			fallthrough
		case '{':
			fallthrough
		case '<':
			s.push(c)
		case ')': // Closing chuck
			fallthrough
		case ']':
			fallthrough
		case '}':
			fallthrough
		case '>':
			opening := s.pop()
			if c != matching[opening] {
				// Discard corrupted line
				return 0
			}
		}
	}

	// Completion
	score := 0
	for !s.isEmpty() {
		closing := matching[s.pop()]
		score = score*5 + completionPoints[closing]
	}
	return score
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var line string
	var err error
	var scores []int
	for line, err = r.ReadString('\n'); err == nil; line, err = r.ReadString('\n') {
		line := line[:len(line)-1] // Removing nl
		if score := autocomplete(line); score != 0 {
			scores = append(scores, score)
		}
	}
	if score := autocomplete(line); score != 0 {
		scores = append(scores, score)
	}

	// Sorting scores & finding the middle score.
	sort.Ints(scores)
	fmt.Println(scores[len(scores)/2])
}
