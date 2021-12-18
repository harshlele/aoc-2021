package main

import (
	utils "aoc-2021/aoc-utils"
	"fmt"
	"math/rand"
)

type Node struct {
	id     int
	val    [2]int
	left   *Node
	right  *Node
	order  [2]string
	parent *Node
	root   *Node
	height int
}

func main() {
	/*
		content, err := ioutil.ReadFile("test-input.txt")
		if err != nil {
			panic(err)
		}

		lines := strings.Split(string(content), "\r\n")


			for _, line := range lines {
			}
	*/
	n := parseLine("[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]")

	utils.X(n)
}

func initNode() Node {
	return Node{rand.Intn(100), [2]int{-1, -1}, nil, nil, [2]string{"", ""}, nil, nil, 0}
}

func parseLine(line string) Node {
	root := initNode()
	curr := &root

	i := 1
	for i < len(line)-1 {
		fmt.Println(i)
		if string(line[i]) == "[" {
			//n := Node{[2]int{-1, -1}, []Node{}, curr, curr.height + 1, &root, rand.Intn(100)}
			n := initNode()
			n.parent = curr
			n.root = &root
			n.height = curr.height + 1

			if curr.left == nil {
				curr.left = &n
				if curr.order[0] == "" {
					curr.order[0] = "node"
				} else {
					curr.order[1] = "node"
				}
				curr = curr.left
			} else if curr.right == nil {
				curr.right = &n
				if curr.order[0] == "" {
					curr.order[0] = "node"
				} else {
					curr.order[1] = "node"
				}
				curr = curr.right
			} else {
				panic(fmt.Sprintf("trying to insert 3 children at i = %d", i))
			}

			i++
		} else if string(line[i]) == "]" {
			curr = curr.parent
			i++
		} else if string(line[i]) == "," {
			i++
		} else {
			num := []byte{}
			for string(line[i]) != "," && string(line[i]) != "]" {
				num = append(num, line[i])
				i++
			}
			val := utils.ToInt(string(num))
			if curr.val[0] == -1 {
				curr.val[0] = val
			} else if curr.val[1] == -1 {
				curr.val[1] = val
			}
			if curr.order[0] == "" {
				curr.order[0] = "int"
			} else {
				curr.order[1] = "int"
			}

			if string(line[i]) == "]" {
				curr = curr.parent
			} else if string(line[i]) != "," {
				panic(fmt.Sprintf("trying to insert 3 ints at i = %d", i))
			}

			i++
		}

	}

	return root
}
