package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strings"
	"aoc-2021/aoc-utils"
)

func main() {

	content, err := ioutil.ReadFile("day7/input.txt")
	if err != nil {
		panic(err)
	}

	s := strings.Split(string(strings.Trim(string(content), "\n")), ",")
	nos := utils.StrToInt(s)

	minDist, max := calcP1(nos)
	fmt.Println(minDist)

	calcP2(nos, max)

}

//for calculating part 1
func calcP1(nos []int) (int, int) {
	arr := nos[:]
	//sort the array
	sort.Ints(arr)

	minDist := math.MaxInt
	a := 0
	b := 0
	//check the sum of the distance for values close to the median
	if len(arr)%2 != 0 {
		a, b = findDist(arr, len(arr)/2), findDist(arr, len(arr)/2-1)
	} else {
		a, b = findDist(arr, len(arr)/2), findDist(arr, len(arr)/2+1)
	}
	if a < b {
		minDist = a
	} else {
		minDist = b
	}

	return minDist, arr[len(arr)-1]
}

//just calculating the normal way lol, median doesnt work for part 2
func calcP2(arr []int, max int) {
	overallMin := math.MaxInt

	for i := 0; i < max; i++ {
		dist := findDistP2(arr, i, overallMin)
		if dist != -1 {
			overallMin = dist
		}
	}

	fmt.Println(overallMin)
}

func findDist(arr []int, index int) int {
	sum := 0
	for idx, _ := range arr {
		sum += int(math.Abs(float64(arr[idx] - arr[index])))
	}
	return sum
}

func findDistP2(arr []int, val int, currMin int) int {
	sum := 0
	for idx, _ := range arr {
		dist := int(math.Abs(float64(arr[idx] - val)))
		sum += (dist * (dist + 1)) / 2
		if sum > currMin {
			return -1
		}
	}
	return sum
}

