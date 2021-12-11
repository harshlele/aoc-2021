package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {

	content, err := ioutil.ReadFile("day10/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\r\n")

	filledStacks := calcP1(lines)
	fmt.Println(calcP2(filledStacks))
}

//calculate the penalty by iterating the stack in reverse
func calcP2(stacks [][]string) int {
	ptList := []int{}
	for _, stack := range stacks {
		points := 0
		for i := len(stack) - 1; i >= 0; i-- {
			a := stack[i]
			points *= 5
			switch a {
			case "(":
				points += 1
			case "[":
				points += 2
			case "{":
				points += 3
			case "<":
				points += 4
			}
		}
		ptList = append(ptList, points)
	}
	sort.Ints(ptList)
	return ptList[len(ptList)/2]
}

//have a stack that pushed opening tags in the stack, then when it comes across a closing tag,
//pop the stack and check if the tags are right (if theyre not right, thats a corrupt string)
//at the end of each line, if the stacks are not full (ie its not corrupt, but it is incomplete),
//store those stacks and return them to use in part 2
func calcP1(lines []string) [][]string {
	totalPts := 0
	inCompStacks := [][]string{}
	symMap := map[string]string{">": "<", "}": "{", "]": "[", ")": "("}
	penaltyMap := map[string]int{">": 25137, "}": 1197, "]": 57, ")": 3}

	for _, line := range lines {
		stack := []string{}
		points := 0
		for i := 0; i < len(line); i++ {
			if strings.Contains("<{[(", string(line[i])) {
				stack = append(stack, string(line[i]))
			} else {
				top := ""
				if len(stack) > 0 {
					top = stack[len(stack)-1]
					stack = stack[:len(stack)-1]
				}
				if top != symMap[string(line[i])] {
					points += penaltyMap[string(line[i])]
				}
			}
		}
		if points > 0 {
			totalPts += points
		} else {
			inCompStacks = append(inCompStacks, stack)
		}

	}
	fmt.Println(totalPts)
	return inCompStacks
}
