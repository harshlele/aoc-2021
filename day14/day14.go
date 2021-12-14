package main

import (
	utils "aoc-2021/aoc-utils"
	"fmt"
	"io/ioutil"
	"math"
	re "regexp"
	"strings"
)

func main() {

	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	exp := re.MustCompile("\\r\\n\\r\\n")

	inputs := exp.Split(string(content), -1)
	template := []byte(inputs[0])
	lines := strings.Split(inputs[1], "\r\n")
	rules := map[string]string{}
	for _, l := range lines {
		k := strings.Split(l, " -> ")
		rules[k[0]] = k[1]
	}

	//fmt.Println(template)
	//fmt.Println(rules)

	calcP1(template, rules, 10)
}

func calcP1(bArr []byte, rules map[string]string, steps int) {
	count := map[byte]int{}
	max := 0
	for _, b := range bArr {
		_, ok := count[b]
		if !ok {
			count[b] = 1
		} else {
			count[b] += 1
			if count[b] > max {
				max = count[b]
			}
		}
	}

	for i := 0; i < steps; i++ {
		j := 0
		for j < len(bArr) {
			val, ok := rules[string(bArr[j:j+2])]
			if ok {
				bArr = utils.InsertAtBArr(bArr, j+1, []byte(val))

				b := []byte(val)[0]
				c, ok := count[b]
				if !ok {
					count[b] = 1
				} else {
					count[b] = c + 1
					if count[b] > max {
						max = count[b]
					}
				}
				j += 2
			} else {
				j += 1
			}
		}
	}

	//fmt.Println(bArr)
	min := math.MaxInt
	for k := range count {
		if min > count[k] {
			min = count[k]
		}
	}
	fmt.Println(max - min)
}
