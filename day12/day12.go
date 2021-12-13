package main

import (
	utils "aoc-2021/aoc-utils"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	content, err := ioutil.ReadFile("day12/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\r\n")
	edges := [][]string{}

	graph := map[string][]string{}

	for _, l := range lines {
		e := strings.Split(l, "-")
		_, ok := graph[e[0]]
		if !ok {
			graph[e[0]] = []string{e[1]}
		} else {
			newArr := append(graph[e[0]], e[1])
			graph[e[0]] = newArr
		}

		_, ok2 := graph[e[1]]
		if !ok2 {
			graph[e[1]] = []string{e[0]}
		} else {
			newArr := append(graph[e[1]], e[0])
			graph[e[1]] = newArr
		}

		edges = append(edges, e)
	}

	//indicates whether ANY lowercase letter has been repeated in the current path(ie stack)
	twiceVisited := false
	fmt.Println(calcRes(graph, "start", []string{}, &twiceVisited))

}

func calcRes(graph map[string][]string, node string, stack []string, p2Case *bool) int {
	//if node is end, go back
	if node == "end" {
		return 1
	}
	sum := 0

	stack = append(stack, node)

	for _, n := range graph[node] {
		//never revisit the start
		if n == "start" {
			continue
		}

		if n != strings.ToLower(n) {
			sum += calcRes(graph, n, stack, p2Case)
		} else {
			//if its a lower case letter, if it has never been visited, just visit it
			//if it has been visited once, and there is no letter in the current path that has been visited twice, visit it
			cnt := utils.CountInArr(stack, n)
			if cnt == 0 {
				sum += calcRes(graph, n, stack, p2Case)
			} else if cnt == 1 && !(*p2Case) {
				*p2Case = true
				sum += calcRes(graph, n, stack, p2Case)
			}
		}
	}

	//IMPORTANT
	//when backtracking, if the current node was the one that got visited twice
	//and thus set p2Case to true, set it to false again when popping the node off the stack
	last := stack[len(stack)-1]
	if last == strings.ToLower(last) && utils.CountInArr(stack, last) == 2 {
		*p2Case = false
	}
	stack = stack[:len(stack)-1]

	return sum
}
