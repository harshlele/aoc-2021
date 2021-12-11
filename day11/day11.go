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

	octo := [][]int{}

	lines := strings.Split(string(content), "\r\n")
	for _, l := range lines {
		nums := strings.Split(l, "")
		octo = append(octo, strToInt(nums))
	}

	calc(octo, 100, false)
}

func calc(arr [][]int, steps int, part2 bool) {
	totalFlashes := 0
	for i := 0; i < steps; i++ {
		flashMap := map[string]bool{}
		for j, _ := range arr {

			for k, _ := range arr[j] {
				checkFlash(arr, j, k, flashMap)
			}
		}

		//have to manually set flashed indexes to 0 for some reason...
		for k := range flashMap {
			sp := strToInt(strings.Split(k, ","))
			arr[sp[0]][sp[1]] = 0
		}
		//for part 2, just check the step at which every octopus flashes all at once
		if len(flashMap) == 100 && part2 {
			fmt.Println("ALL FLASH AT ", i+1)
			break
		}
		totalFlashes += (len(flashMap))
	}
	fmt.Println(totalFlashes)
}

func checkFlash(arr [][]int, i, j int, flashes map[string]bool) {
	if i < 0 || i >= len(arr) || j < 0 || j >= len(arr[0]) {
		return
	}
	arr[i][j] += 1
	if arr[i][j] > 9 {
		_, ok := flashes[pointKey([]int{i, j})]
		if !ok {
			flashes[pointKey([]int{i, j})] = true
			//increment adjacent indexes
			checkFlash(arr, i-1, j-1, flashes)
			checkFlash(arr, i-1, j, flashes)
			checkFlash(arr, i-1, j+1, flashes)
			checkFlash(arr, i, j-1, flashes)
			checkFlash(arr, i, j+1, flashes)
			checkFlash(arr, i+1, j-1, flashes)
			checkFlash(arr, i+1, j, flashes)
			checkFlash(arr, i+1, j+1, flashes)

			//setting it to 0 at the end should work, but doesnt for some reason lol
			//arr[i][j] = 0
		}

	}
}

//just converts string array to int array
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
