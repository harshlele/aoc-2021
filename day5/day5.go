package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	points := [][][]int{}
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	//split lines
	lines := strings.Split(string(content), "\r\n")
	//convert lines to 3D array - (array of lines, each line containing 2 points, each point containing x and y coordinates)
	for i := 0; i < len(lines); i++ {
		p := strings.Split(lines[i], " -> ")

		c1 := strToInt(strings.Split(p[0], ","))
		c2 := strToInt(strings.Split(p[1], ","))
		pair := [][]int{}
		pair = append(pair, c1)
		pair = append(pair, c2)
		points = append(points, pair)

	}

	calcP1(points)

}

//part 1
func calcP1(pointPairs [][][]int) {
	coveredPoints := map[string]int{}

	for i := 0; i < len(pointPairs); i++ {
		pair := pointPairs[i]
		//if not horizontal/vertical, skip
		dir := isHorOrVert(pair)
		if dir == -1 {
			continue
		}
		//mark the 2 points as covered
		p1 := pointKey(pair[0])
		p2 := pointKey(pair[1])
		insertIntoPoints(coveredPoints, p1)
		insertIntoPoints(coveredPoints, p2)

		//get init and upto for looping
		init, upto := sort2(pair[0][dir], pair[1][dir])
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
			key := pointKey(point)
			insertIntoPoints(coveredPoints, key)

			j++
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
//else return -1
func isHorOrVert(line [][]int) int {
	if line[0][0] == line[1][0] {
		return 1
	} else if line[0][1] == line[1][1] {
		return 0
	} else {
		return -1
	}
}

//just converts string array to int array
func strToInt(arr []string) []int {
	a := []int{}

	for i := 0; i < len(arr); i++ {
		I, err := strconv.Atoi(arr[i])
		if err != nil {
			panic(err)
		}
		a = append(a, I)
	}

	return a
}

func pointKey(point []int) string {
	return fmt.Sprintf("%d,%d", point[0], point[1])
}

//returns sorted order
func sort2(a int, b int) (int, int) {
	if a < b {
		return a, b
	} else {
		return b, a
	}
}

//for unused variables lol
func x(X ...interface{}) {}
