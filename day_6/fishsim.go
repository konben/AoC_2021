package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const noOfDays = 256

func main() {
	// Read initial population
	r := bufio.NewReader(os.Stdin)
	input, _ := r.ReadString('\n')
	timers := strings.Split(input, ",")
	fishes := map[int]int{}
	for _, t := range timers {
		timer, _ := strconv.ParseInt(t, 10, 64)
		fishes[int(timer)]++
	}

	// Simulate reproduction
	for i := 0; i < noOfDays; i++ {
		next := map[int]int{}
		for timer, noOfFish := range fishes {
			if timer == 0 {
				next[8] += noOfFish
				next[6] += noOfFish
				continue
			}
			next[timer-1] += noOfFish
		}
		fishes = next
	}

	// Count fish
	totalFish := 0
	for _, noOfFish := range fishes {
		totalFish += noOfFish
	}
	fmt.Println(totalFish)
}
