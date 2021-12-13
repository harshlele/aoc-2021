package main

import (
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

	//for counting paths
	ct := 0
	p := &ct

	//for counting the exceptional small cave
	ec := 0
	p2 := &ct

	calcRes(graph, "start", map[string]int{}, []string{}, &p, &p2, "b")
	fmt.Println(ct)
	fmt.Println(ec)
}

func calcRes(graph map[string][]string, node string, visited map[string]int, stack []string, counter **int, exceptionalCt **int, except string) {

	visited[node] = 1
	stack = append(stack, node)
	//fmt.Println(stack)

	pt := *counter

	for _, n := range graph[node] {
		_, ok := visited[n]
		if (!ok || n != strings.ToLower(n)) && node != "end" {

			calcRes(graph, n, visited, stack, counter, exceptionalCt, except)
		}
		if n == "end" {
			(*pt) += 1
		}
	}
	stack = stack[:len(stack)-1]
	delete(visited, node)
}
