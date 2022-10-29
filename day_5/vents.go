package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type coordinate [2]int

func (c coordinate) x() int {
	return c[0]
}

func (c coordinate) y() int {
	return c[1]
}

type ventLine []coordinate

func parseLine(line string) ventLine {
	// Find start/end coordinate
	vecs := strings.Split(line, " -> ")
	start := parseCoordinate(vecs[0])
	end := parseCoordinate(vecs[1])

	// Compute line
	var vents ventLine
	xDiff, yDiff := end.x()-start.x(), end.y()-start.y()
	xStep, yStep := 0, 0
	steps := int(math.Max(math.Abs(float64(xDiff)), math.Abs(float64(yDiff)))) + 1
	if xDiff != 0 {
		xStep = xDiff / int(math.Abs(float64(xDiff)))
	}
	if yDiff != 0 {
		yStep = yDiff / int(math.Abs(float64(yDiff)))
	}
	x, y := start.x(), start.y()
	for i := 0; i < steps; i++ {
		vents = append(vents, coordinate{x, y})
		x += xStep
		y += yStep
	}
	return vents
}

func parseCoordinate(coord string) coordinate {
	coords := strings.Split(coord, ",")
	x1, _ := strconv.ParseInt(coords[0], 10, 64)
	x2, _ := strconv.ParseInt(coords[1], 10, 64)
	return coordinate{int(x1), int(x2)}
}

func main() {
	// Read lines
	r := bufio.NewReader(os.Stdin)
	var line string
	var err error
	var ventLines []ventLine
	for line, err = r.ReadString('\n'); err == nil; line, err = r.ReadString('\n') {
		ventLines = append(ventLines, parseLine(line[:len(line)-1]))
	}
	ventLines = append(ventLines, parseLine(line))

	// Visualise example
	/*
		vents := [10][10]int{}
		for _, line := range ventLines {
			for _, coord := range line {
				x, y := coord[0], coord[1]
				fmt.Println(x, y)
				vents[x][y]++
			}
		}
		for y := 0; y < 10; y++ {
			for x := 0; x < 10; x++ {
				ventCount := vents[x][y]
				if ventCount == 0 {
					fmt.Print(".")
				} else {
					fmt.Print(ventCount)
				}
			}
			fmt.Println("")
		}
	*/

	// Count overlaps
	vents := make(map[coordinate]int)
	overlaps := 0
	for _, line := range ventLines {
		for _, coord := range line {
			vents[coord]++

			// Did exactly two lines overlap?
			if vents[coord] == 2 {
				overlaps++
			}
		}
	}
	fmt.Println("Overlaps:", overlaps)
}
