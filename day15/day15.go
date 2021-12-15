package main

import (
	utils "aoc-2021/aoc-utils"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	vQ := PriorityQ{map[string][]int{}, ""}
	uQ := PriorityQ{map[string][]int{}, ""}

	lines := strings.Split(string(content), "\r\n")

	for i, l := range lines {
		ints := utils.StrToInt(strings.Split(l, ""))
		for j, I := range ints {
			if i == 0 && j == 0 {
				uQ.Push(i, j, []int{I, 0})
			} else {
				uQ.Push(i, j, []int{I, 100000})
			}

		}
	}

	getShortestPath2(vQ, uQ)

	//set og cost to 0
	//getShortestPath(rMap, visited, 0, 0)
}

func getShortestPath2(vQ, uQ PriorityQ) {

	isValid := func(i, j int) bool {
		if i < 0 || i >= 100 || j < 0 || j >= 100 {
			return false
		}
		return true
	}

	updateDist := func(p0, p1 int, par []int) {
		if isValid(p0, p1) {
			key := utils.PointKey([]int{p0, p1})
			if uQ.isInMap(key) {
				val := uQ.Map[key]
				if val[1] > val[0]+par[1] {
					uQ.UpdateAtKey(p0, p1, []int{val[0], val[0] + par[1]})
				}

			} else {
				if vQ.isInMap(key) {
					val := vQ.Map[key]
					if val[1] > val[0]+par[1] {
						vQ.UpdateAtKey(p0, p1, []int{val[0], val[0] + par[1]})
					}

				}

			}

		}
	}

	initL := len(uQ.Map)
	fmt.Println(initL)
	for true {
		min := utils.StrToInt(strings.Split(uQ.min, ","))

		val := uQ.Remove(min[0], min[1])
		vQ.Push(min[0], min[1], val)

		t0, t1, b0, b1, l0, l1, r0, r1 := min[0]-1, min[1], min[0]+1, min[1], min[0], min[1]-1, min[0], min[1]+1

		updateDist(t0, t1, val)
		updateDist(b0, b1, val)
		updateDist(l0, l1, val)
		updateDist(r0, r1, val)

		if len(vQ.Map) == initL {
			fmt.Println(vQ.Map["99,99"])
			break
		}
	}
}

func getShortestPath(graph [][][]int, visitedPaths map[string]bool, i, j int) {

	visitedPaths[utils.PointKey([]int{i, j})] = true

	updateAndTraverse := func(pX, pY int) {

		curr := graph[i][j]
		adj := graph[pX][pY]

		if adj[1] > curr[1]+adj[0] {
			graph[pX][pY][1] = curr[1] + adj[0]
		}

	}

	u0, u1, d0, d1, l0, l1, r0, r1 := i-1, j, i+1, j, i, j-1, i, j+1

	updateAndTraverse(u0, u1)
	updateAndTraverse(d0, d1)
	updateAndTraverse(l0, l1)
	updateAndTraverse(r0, r1)

}
