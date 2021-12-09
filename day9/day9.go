package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

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

	calcP1(nos)
}

func calcP1(nos [][]int) {
	mins := 0
	for i, val := range nos {
		for j, num := range val {
			h1, h2, v1, v2 := 0, 0, 0, 0

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
			}
		}
	}
	fmt.Println(mins)
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
