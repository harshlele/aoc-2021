package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\r\n")
	fmt.Println(calcP1(lines))

}

func calcP1(lines []string) int {
	sum := 0

	for _, line := range lines {
		output := strings.Split(line, "|")
		words := strings.Split(output[1], " ")

		lMap := map[int]bool{2: true, 4: true, 3: true, 7: true}

		for _, word := range words {
			_, ok := lMap[len(word)]
			if ok {
				sum += 1
			}
		}
	}

	return sum
}
