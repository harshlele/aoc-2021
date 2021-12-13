package utils

import (
	"fmt"
	"strconv"
)

//just converts string array to int array
func StrToInt(arr []string) []int {
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

func ToInt(a string) int {
	I, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}
	return I
}

func PointKey(point []int) string {
	return fmt.Sprintf("%d,%d", point[0], point[1])
}

//for unused variables lol
func X(x ...interface{}) {}

//returns sorted order
func Sort2(a int, b int) (int, int) {
	if a < b {
		return a, b
	} else {
		return b, a
	}
}

func CountInArr(arr []string, str string) int {
	ct := 0
	for _, val := range arr {
		if val == str {
			ct++
		}
	}
	return ct
}
