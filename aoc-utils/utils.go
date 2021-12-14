package utils

import (
	"bytes"
	"fmt"
	"strconv"
)

//just converts string array to int array
func StrToInt(arr []string) []int {
	a := []int{}
	for _, val := range arr {
		a = append(a, ToInt(val))
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

//just a simple string derived from 2 ints to use as a key in maps etc
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

//insert bytes in byte array
func InsertAtBArr(arr []byte, position int, newBytes []byte) []byte {

	newArr := bytes.Buffer{}
	newArr.Write(arr[:position])
	newArr.Write(newBytes)
	newArr.Write(arr[position:])

	return newArr.Bytes()
}
