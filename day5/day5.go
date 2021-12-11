package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"aoc-2021/aoc-utils"
)

func main() {

	points := [][][]int{}

	//set true to get the answer for part 2!
	part2 := false

	content, err := ioutil.ReadFile("day5/test-input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\r\n")
	//convert lines to 3D array - (array of lines, each line containing 2 points, each point containing x and y coordinates)
	for i := 0; i < len(lines); i++ {
		p := strings.Split(lines[i], " -> ")

		c1 := utils.StrToInt(strings.Split(p[0], ","))
		c2 := utils.StrToInt(strings.Split(p[1], ","))
		pair := [][]int{}
		pair = append(pair, c1)
		pair = append(pair, c2)
		points = append(points, pair)

	}

	calcRes(points, part2)

}

func calcRes(pointPairs [][][]int, part2 bool) {
	coveredPoints := map[string]int{}

	for i := 0; i < len(pointPairs); i++ {
		pair := pointPairs[i]

		//if not horizontal/vertical and part 1, skip
		dir := isHorOrVert(pair)
		if dir == -1 && !part2 {
			continue
		}

		//mark the 2 end points as covered
		p1, p2 := utils.PointKey(pair[0]), utils.PointKey(pair[1])
		insertIntoPoints(coveredPoints, p1)
		insertIntoPoints(coveredPoints, p2)

		if !part2 || (part2 && dir != -1) {
			init, upto := utils.Sort2(pair[0][dir], pair[1][dir])
			j := init + 1

			for j < upto {

				point := []int{}
				if dir == 0 {
					point = append(point, j)
					point = append(point, pair[0][1])
				} else {
					point = append(point, pair[0][0])
					point = append(point, j)
				}
				key := utils.PointKey(point)
				insertIntoPoints(coveredPoints, key)

				j++
			}
		} else {
			x1, x2 := utils.Sort2(pair[0][0], pair[1][0])
			y1, y2 := utils.Sort2(pair[0][1], pair[1][1])

			for i := x1 + 1; i < x2; i++ {
				for j := y1 + 1; j < y2; j++ {
					if checkPointInLine(pair, []int{i, j}) {
						key := utils.PointKey([]int{i, j})
						insertIntoPoints(coveredPoints, key)
					}
				}
			}
		}
	}

	//count twice-covered points
	ct := 0
	for _, val := range coveredPoints {
		if val >= 2 {
			ct++
		}
	}
	fmt.Println(ct)
}

func insertIntoPoints(Map map[string]int, key string) {
	_, ok := Map[key]
	if !ok {
		Map[key] = 1
	} else {
		Map[key] += 1
	}
}

//if its a horizontal line(ie y coords have to be checked), return 1(index of y co-ordinate for each point)
//if its a vertical line(ie x coords have to be checked), return 0(index of y co-ordinate for each point)
//else return -1 (ie diagonals for part 2)
func isHorOrVert(line [][]int) int {
	if line[0][0] == line[1][0] {
		return 1
	} else if line[0][1] == line[1][1] {
		return 0
	} else {
		return -1
	}
}



//checks whether a point is in a line using the (x-x1)/(x1-x2) = (y-y1)/(y1-y2) equation
func checkPointInLine(line [][]int, point []int) bool {
	var x, y float32
	x = (float32(point[0]) - float32(line[0][0])) / (float32(line[0][0]) - float32(line[1][0]))
	y = (float32(point[1]) - float32(line[0][1])) / (float32(line[0][1]) - float32(line[1][1]))

	return x == y
}

