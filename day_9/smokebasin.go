package main

import (
	"bufio"
	"fmt"
	"os"
)

type heightmap [][]byte

func (h heightmap) heightWidth() (int, int) {
	height := len(h)
	if height == 0 {
		return 0, 0
	}
	width := len(h[0])
	return height, width
}

// Returns the coordinate (y, x), if (y, x) is out of bounds 0xff is returned.
func (h heightmap) get(y, x int) byte {
	height, width := h.heightWidth()
	if y < 0 || y >= height || x < 0 || x >= width {
		return 0xff
	}
	return h[y][x]
}

// Returns the neighbours of coordinate (y, x).
func (h heightmap) neighbours(y, x int) [4]byte {
	return [4]byte{h.get(y-1, x), h.get(y+1, x), h.get(y, x-1), h.get(y, x+1)}
}

// Computes the risk level of coordinate (y, x), the risk level of points that are not low points is 0.
func (h heightmap) riskLevel(y, x int) int {
	height := h.get(y, x)
	neighbours := h.neighbours(y, x)
	for _, n := range neighbours {
		if height >= n {
			// (y, x) is not a low point.
			return 0
		}
	}
	return int(height-'0') + 1
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var line string
	var err error
	var hmap heightmap
	for line, err = r.ReadString('\n'); err == nil; line, err = r.ReadString('\n') {
		line := line[:len(line)-1] // Removing nl
		hmap = append(hmap, []byte(line))
	}
	hmap = append(hmap, []byte(line))

	// Compute risk levels of low points.
	riskSum := 0
	height, width := hmap.heightWidth()
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			riskSum += hmap.riskLevel(y, x)
		}
	}
	fmt.Println("Sum:", riskSum)
}
