package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"aoc-2021/aoc-utils"
)

func main() {

	content, err := ioutil.ReadFile("day6/input.txt")
	if err != nil {
		panic(err)
	}

	s := strings.Split(string(strings.Trim(string(content), "\n")), ",")

	nos := utils.StrToInt(s)

	days := 256
	utils.X(days)
	calcP2(nos, days)
}

//dumb approach lol
func calcP1(nos []int, days int) {
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

func calcP2(nos []int, days int) {
	sum := 0
	dayMap := map[string]int{}
	for _, val := range nos {
		sum += getFishForDays(val, days, dayMap)
		//fmt.Println("done with ", idx)
	}

	fmt.Println(sum)
}

//uses recursion
func getFishForDays(fish int, days int, dayMap map[string]int) int {
	//if more time for days, no time to split
	if fish >= days {
		return 1
	}

	//fish is 0, return sum of results for 6 and 8, ie split into 2 fish (check map first for both cases)
	//else return results for 0,days - fish (ie skip forward until a split event happens, again check map first)
	if fish == 0 {
		k1, k2 := fmt.Sprintf("6,%d", days-1), fmt.Sprintf("8,%d", days-1)
		c1, c2 := 0, 0

		v1, ok1 := dayMap[k1]
		if !ok1 {
			c1 = getFishForDays(6, days-1, dayMap)
			dayMap[k1] = c1
		} else {
			c1 = v1
		}

		v2, ok2 := dayMap[k2]
		if !ok2 {
			c2 = getFishForDays(8, days-1, dayMap)
			dayMap[k2] = c2
		} else {
			c2 = v2
		}
		return c1 + c2

	} else {
		key := fmt.Sprintf("0,%d", days-fish)
		val, ok := dayMap[key]

		if !ok {
			days := getFishForDays(0, days-fish, dayMap)
			dayMap[key] = days
			return days
		} else {
			return val
		}
	}
}
