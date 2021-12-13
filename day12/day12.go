package main

import (
	utils "aoc-2021/aoc-utils"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	content, err := ioutil.ReadFile("day12/test-input.txt")
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

	twiceVisited := false
	utils.X(twiceVisited)
	fmt.Println(calcRes(graph, "start", []string{}))

}

func calcRes(graph map[string][]string, node string, stack []string) int {

	if node == "end" {
		return 1
	}
	sum := 0
	stack = append(stack, node)

	for _, n := range graph[node] {
		if n != strings.ToLower(n) {
			sum += calcRes(graph, n, stack)
		} else {
			if utils.CountInArr(stack, n) == 0 {
				sum += calcRes(graph, n, stack)
			}
		}
	}
	stack = stack[:len(stack)-1]
	return sum
}
