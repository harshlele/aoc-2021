package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type Basin struct {
	min    []int
	points map[string]bool
}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\r\n")

	nos := [][]int{}

	for _, val := range lines {
		numLine := strToInt(strings.Split(val, ""))
		nos = append(nos, numLine)
	}

	lows := calcP1(nos)

	calcP2(nos, lows)
}

//depth-first traversal starting from the low point
func calcP2(nos [][]int, lows []Basin) {
	sizes := []int{}
	for _, val := range lows {
		sizes = append(sizes, dfsCount(nos, val.min, val.points))
	}
	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})

	fmt.Println(sizes[0] * sizes[1] * sizes[2])
}

func dfsCount(nos [][]int, point []int, points map[string]bool) int {
	i, j := point[0], point[1]
	//check for out of bounds
	if i < 0 || i >= len(nos) || j < 0 || j >= len(nos[0]) {
		return 0
	}
	if nos[i][j] == 9 {
		return 0
	}

	_, ok := points[pointKey(point)]
	//if the point has already been traversed, skip it
	if ok {
		return 0
	} else {
		points[pointKey(point)] = true
		l, r, t, b := []int{i, j - 1}, []int{i, j + 1}, []int{i - 1, j}, []int{i + 1, j}

		//apparently it produces the right answer without checking if adjacent numbers are actually greater
		//(ie if its "flowing downward")
		//but i added it anyway lol
		if l[1] >= 0 {
			if nos[i][j] >= nos[l[0]][l[1]] {
				l = []int{-1, j}
			}
		}
		if r[1] < len(nos[0]) {
			if nos[i][j] >= nos[r[0]][r[1]] {
				r = []int{i, -1}
			}
		}
		if t[0] >= 0 {
			if nos[i][j] >= nos[t[0]][t[1]] {
				t = []int{-1, j}
			}
		}
		if b[0] < len(nos) {
			if nos[i][j] >= nos[b[0]][b[1]] {
				b = []int{-1, j}
			}
		}

		return 1 + dfsCount(nos, l, points) + dfsCount(nos, r, points) + dfsCount(nos, t, points) + dfsCount(nos, b, points)
	}
}

//part 1
//returns the minimums for part 2
func calcP1(nos [][]int) []Basin {
	mins := 0
	lows := []Basin{}
	for i, val := range nos {
		for j, num := range val {
			h1, h2, v1, v2 := 0, 0, 0, 0

			//check for out of bounds indexes
			if j == 0 {
				h1, h2 = 1, 1
			} else if j == len(val)-1 {
				h1, h2 = len(val)-2, len(val)-2
			} else {
				h1, h2 = j-1, j+1
			}

			if i == 0 {
				v1, v2 = 1, 1
			} else if i == len(nos)-1 {
				v1, v2 = len(nos)-2, len(nos)-2
			} else {
				v1, v2 = i-1, i+1
			}

			if num < val[h1] && num < val[h2] && num < nos[v1][j] && num < nos[v2][j] {
				mins += num + 1
				lows = append(lows, Basin{[]int{i, j}, map[string]bool{}})
			}
		}
	}
	fmt.Println(mins)
	return lows
}

func strToInt(arr []string) []int {
	a := []int{}

	for _, val := range arr {
		I, err := strconv.Atoi(val)
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
