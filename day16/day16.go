package main

import (
	utils "aoc-2021/aoc-utils"
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	hex := string(content)
	//utils.X(hex)
	inp := hexToBin(hex)
	fmt.Println(strings.Join(inp, ""))
	fmt.Println(getPacketVerSum(inp, 0))
}

func getPacketVerSum(binary []string, start int) (int, int) {
	ver := binStrToDecInt(binary[start : start+3])
	id := binStrToDecInt(binary[start+3 : start+6])

	if id == 4 {
		upto := start + 6
		for binary[upto] != "0" {
			upto += 5
		}
		upto += 4
		return ver, upto
	} else {
		sum := ver
		fUpto := 0
		if binary[start+6] == "0" {
			l := binStrToDecInt(binary[start+7 : start+7+15])
			upto := start + 7 + 15
			curr := upto
			for upto < curr+l {
				a, b := getPacketVerSum(binary, upto)
				sum += a
				upto = b + 1
			}
			fUpto = upto - 1
		} else if binary[start+6] == "1" {
			l := binStrToDecInt(binary[start+7 : start+7+11])
			i := 0
			upto := start + 7 + 11
			for i < l {
				a, b := getPacketVerSum(binary, upto)
				sum += a
				upto = b + 1
				i++
			}
			fUpto = upto - 1
		}

		return sum, fUpto
	}
}

func hexToBin(hex string) []string {
	binMap := map[string][4]string{
		"0": {"0", "0", "0", "0"},
		"1": {"0", "0", "0", "1"},
		"2": {"0", "0", "1", "0"},
		"3": {"0", "0", "1", "1"},
		"4": {"0", "1", "0", "0"},
		"5": {"0", "1", "0", "1"},
		"6": {"0", "1", "1", "0"},
		"7": {"0", "1", "1", "1"},
		"8": {"1", "0", "0", "0"},
		"9": {"1", "0", "0", "1"},
		"A": {"1", "0", "1", "0"},
		"B": {"1", "0", "1", "1"},
		"C": {"1", "1", "0", "0"},
		"D": {"1", "1", "0", "1"},
		"E": {"1", "1", "1", "0"},
		"F": {"1", "1", "1", "1"},
	}

	binary := make([]string, len(hex)*4)

	for i := 0; i < len(hex); i++ {
		b := binMap[string(hex[i])]
		binary[4*i], binary[4*i+1], binary[4*i+2], binary[4*i+3] = b[0], b[1], b[2], b[3]
	}

	return binary
}

func binStrToDecInt(binary []string) int {
	dec := 0
	for i := len(binary) - 1; i >= 0; i-- {
		digit := utils.ToInt(string(binary[i]))
		pos := float64(len(binary) - i - 1)

		dec += digit * int(math.Pow(2, pos))
	}

	return dec
}
