package main

import (
	utils "aoc-2021/aoc-utils"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func main() {

	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	vQ := PriorityQ{make(map[string][]int, 250100), ""}
	uQ := PriorityQ{make(map[string][]int, 250100), ""}

	lines := strings.Split(string(content), "\r\n")

	for i, l := range lines {
		ints := utils.StrToInt(strings.Split(l, ""))
		for j, I := range ints {

			for k := 0; k < 5; k++ {
				for m := 0; m < 5; m++ {
					dist := 100000
					if i == 0 && j == 0 && k == 0 && m == 0 {
						dist = 0
					}
					val := I + k + m
					if val > 9 {
						val -= 9
					}

					uQ.Push(i+(k*100), j+(m*100), []int{val, dist})

				}
			}

		}

	}

	getShortestPath2(vQ, uQ)

}

func getShortestPath2(vQ, uQ PriorityQ) {

	isValid := func(i, j int) bool {
		if i < 0 || i >= 500 || j < 0 || j >= 500 {
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
	t := time.Now()
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
			fmt.Println(vQ.Map["499,499"])
			break
		} else if len(vQ.Map)%1000 == 0 {
			fmt.Println(len(vQ.Map), "visited, ", time.Since(t), "since last update")
			t = time.Now()
		}
	}
}
