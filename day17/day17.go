package main

import (
	utils "aoc-2021/aoc-utils"
	"fmt"
	"math"
)

func main() {
	//smol input today :D
	box := [2][2]int{{20, 30}, {-10, -5}}

	//just test for a wide range of values of vx and vy
	vX, vY := 1, 1000
	maxHeight := 0
	//keep map for distinct values
	vMap := map[string]bool{}

	i, j := vX, vY

	for i < 1000 {
		j = vY
		for j > -1000 {

			r, h := runSim(i, j, box)
			if r == true {
				vMap[utils.PointKey([]int{i, j})] = true
				if h > maxHeight {
					maxHeight = h
				}
			}

			j--
		}

		i++
	}

	fmt.Println(maxHeight)
	fmt.Println(len(vMap))
}

func runSim(vX, vY int, box [2][2]int) (bool, int) {
	x, y, Vx, Vy := 0, 0, vX, vY

	maxY := math.MinInt

	for true {

		x += Vx
		y += Vy
		if y > maxY {
			maxY = y
		}

		if Vx > 0 {
			Vx -= 1
		} else if Vx < 0 {
			Vx += 1
		}

		Vy -= 1

		if isInBox(box, x, y) {
			return true, maxY
		} else if isPastBox(box, x, y) {
			return false, maxY
		}
	}

	return false, maxY
}

//check if the probe has overshot the target
func isPastBox(box [2][2]int, x, y int) bool {
	if x > box[0][1] || y < box[1][0] {
		return true
	}
	return false
}

//check if the probe is in the target
func isInBox(box [2][2]int, x, y int) bool {
	if x >= box[0][0] && x <= box[0][1] && y >= box[1][0] && y <= box[1][1] {
		return true
	}
	return false
}
