package main

import (
	"bufio"
	"fmt"
	"os"
)

var matching = map[byte]byte{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

var points = map[byte]int{
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
				return points[c]
			}
		}
	}
	return 0
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var line string
	var err error
	score := 0
	for line, err = r.ReadString('\n'); err == nil; line, err = r.ReadString('\n') {
		line := line[:len(line)-1] // Removing nl
		score += syntaxCheck(line)
	}
	score += syntaxCheck(line)
	fmt.Println(score)
}
