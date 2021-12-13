package main

import (
	utils "aoc-2021/aoc-utils"
	"fmt"
	"io/ioutil"
	re "regexp"
	"strings"
)

func main() {

	content, err := ioutil.ReadFile("day13/input.txt")
	if err != nil {
		panic(err)
	}

	exp := re.MustCompile("\\r\\n\\r\\n")

	inputs := exp.Split(string(content), -1)
	lines := strings.Split(inputs[0], "\r\n")
	mArr := strings.Split(inputs[1], "\r\n")
	coordinates := [][]int{}
	moves := [][]string{}

	for _, l := range lines {
		c := utils.StrToInt(strings.Split(l, ","))
		coordinates = append(coordinates, c)
	}

	for _, m := range mArr {
		s := strings.Split(m, " ")
		moves = append(moves, strings.Split(s[2], "="))
	}

	utils.X(mArr)

	dots := fold(coordinates, moves)

	fmt.Println(len(dots))
	print(dots)

}

//part 2 - print horizontally lol
func print(dots map[string]bool) {
	str := ""
	maxX := 0
	maxY := 0
	//get the max range for the dots
	for k := range dots {
		coords := utils.StrToInt(strings.Split(k, ","))
		if coords[0] > maxX {
			maxX = coords[0] + 5
		}
		if coords[1] > maxY {
			maxY = coords[1] + 5
		}
	}

	//print Y horizontally and X vertically cos thats how the letters appear in the pattern lol
	for i := 0; i < maxY; i++ {
		for j := 0; j < maxX; j++ {
			_, ok := dots[utils.PointKey([]int{j, i})]
			if !ok {
				str += "."
			} else {
				str += "#"
			}

		}
		str += "\n"
	}

	fmt.Println(str)
}

//for part 1
func fold(coordinates [][]int, moves [][]string) map[string]bool {
	//make a map so each point is covered only once
	pointMap := map[string]bool{}

	for _, M := range moves {
		for i, pair := range coordinates {

			if string(M[0]) == "y" {

				if pair[1] < utils.ToInt(string(M[1])) {
					pointMap[utils.PointKey(coordinates[i])] = true
					continue
				}

				//if the point is to be folded(ie shifted up), remove the previous entry from the map if it exists
				_, ok := pointMap[utils.PointKey(pair)]
				if ok {
					delete(pointMap, utils.PointKey(pair))
				}

				//shift up the point
				coordinates[i][1] = utils.ToInt(string(M[1])) - (pair[1] - utils.ToInt(string(M[1])))

			} else if string(M[0]) == "x" {
				if pair[0] < utils.ToInt(string(M[1])) {
					pointMap[utils.PointKey(coordinates[i])] = true
					continue
				}

				//if the point is to be folded(ie shifted left), remove the previous entry from the map if it exists
				_, ok := pointMap[utils.PointKey(pair)]
				if ok {
					delete(pointMap, utils.PointKey(pair))
				}
				//shift left the point
				coordinates[i][0] = utils.ToInt(string(M[1])) - (pair[0] - utils.ToInt(string(M[1])))
			}

			//store the new coordinates in the map
			pointMap[utils.PointKey(coordinates[i])] = true
		}

	}

	return pointMap
}
