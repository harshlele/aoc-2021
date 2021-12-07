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

	s := strings.Split(string(strings.Trim(string(content), "\n")), ",")

	nos := strToInt(s)

	//fmt.Println(nos)

	days := 256
	i := 0

	for i < days {
		toAppend := []int{}

		for idx, fish := range nos {
			if fish == 0 {
				nos[idx] = 6
				toAppend = append(toAppend, 8)
			} else {
				nos[idx]--
			}
		}

		if len(toAppend) > 0 {
			for _, a := range toAppend {
				nos = append(nos, a)
			}
		}
		i++
	}

	fmt.Println(len(nos))
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

//for unused variables lol
func x(X ...interface{}) {}
