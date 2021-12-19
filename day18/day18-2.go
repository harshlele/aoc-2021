package main

import (
	utils "aoc-2021/aoc-utils"
	"fmt"
	"math/rand"
)

type Node struct {
	id, val, height     int
	left, right, parent *Node
}

func main() {
	s := "[[[[4,3],4],4],[7,[[8,4],9]]]"
	a := parseLine(s)
	b := parseLine("[1,1]")
	newR := addNode(a, b)
	utils.X(newR)
}

func initNode() Node {
	return Node{rand.Intn(1000), -1, 0, nil, nil, nil}
}

func addNode(r1, r2 *Node) *Node {
	newRoot := initNode()
	newRoot.left = r1
	newRoot.right = r2
	updateHeights(&newRoot)
	fmt.Println()
	reduce(&newRoot)
	inOrderTraverse(&newRoot)

	return &newRoot
}

func inOrderTraverse(node *Node) {
	if node.val == -1 {
		fmt.Print("[")
	}
	if node.left != nil {
		inOrderTraverse(node.left)
	}
	if node.val != -1 {
		fmt.Print("{", node.val, "}")
	}
	if node.right != nil {
		inOrderTraverse(node.right)
	}
	if node.val == -1 {
		fmt.Print("]")
	}

}

func getSuccessor(node *Node) (bool, *Node) {
	curr := node

	for curr.parent != nil {
		curr = curr.parent
		if curr.right != nil {
			if curr.right.val != -1 {
				break
			}
		}
	}

	if curr.parent == nil {
		return false, nil
	}

	return true, curr.right
}

func getPredecessor(node *Node) (bool, *Node) {
	curr := node

	for curr.parent != nil {
		curr = curr.parent
		if curr.left != nil {
			if curr.left.val != -1 {
				break
			}
		}
	}

	if curr.parent == nil {
		return false, nil
	}

	return true, curr.left
}

func explode(node *Node) {

	var spPred, spSucc *Node

	if node.left != nil {
		v1 := node.left.val
		if v1 != -1 {
			s1, pred := getPredecessor(node)
			if s1 {
				pred.val += v1
				if pred.val > 9 {
					spPred = pred
				}
			}
		}
	}

	if node.right != nil {
		v2 := node.right.val
		if v2 != -1 {
			s2, succ := getSuccessor(node)
			if s2 {
				succ.val += v2
				if succ.val > 9 {
					spSucc = succ
				}
			}
		}
	}

	p := node.parent
	if node == node.parent.left {
		n := initNode()
		p.left = &n
		p.left.parent = p
		p.left.val = 0
		p.left.height = p.height + 1
	} else {
		n := initNode()
		p.right = &n
		p.right.parent = p
		p.right.val = 0
		p.right.height = p.height + 1
	}

	if spPred != nil {
		split(spPred)
	}
	if spSucc != nil {
		split(spSucc)
	}
}

func split(node *Node) {
	p := node.parent
	if p == nil {
		return
	}
	v1, v2 := 0, 0
	if node.val%2 == 0 {
		v1, v2 = node.val/2, node.val/2
	} else {
		v1, v2 = node.val/2, (node.val+1)/2
	}

	if p.left == node {
		n := initNode()
		p.left = &n
		p.left.height = p.height + 1
		p.left.parent = p

		newLeft := initNode()
		p.left.left = &newLeft
		p.left.left.val = v1
		p.left.left.height = p.left.height + 1
		p.left.left.parent = p.left

		newRight := initNode()
		p.left.right = &newRight
		p.left.right.val = v2
		p.left.right.height = p.left.height + 1
		p.left.right.parent = p.left

	} else {
		n := initNode()
		p.right = &n
		p.right.height = p.height + 1
		p.right.parent = p

		newLeft := initNode()
		p.right.left = &newLeft
		p.right.left.val = v1
		p.right.left.height = p.right.height + 1
		p.right.left.parent = p.right

		newRight := initNode()
		p.right.right = &newRight
		p.right.right.val = v2
		p.right.right.height = p.right.height + 1
		p.right.right.parent = p.right

	}

	if p.left.height >= 4 {
		explode(p.left)
	}
	if p.right.height >= 4 {
		explode(p.right)
	}

}

func reduce(node *Node) {

	if node.height < 4 {
		if node.left != nil {
			reduce(node.left)
		}
		if node.right != nil {
			reduce(node.right)
		}
	} else {
		if node.left == nil && node.right == nil {
			return
		}
		explode(node)
	}

}

func updateHeights(node *Node) {

	if node.left != nil {
		node.left.height = node.height + 1
		updateHeights(node.left)
	}
	if node.right != nil {
		node.right.height = node.height + 1
		updateHeights(node.right)
	}
}

func parseLine(line string) *Node {
	root := initNode()
	curr := &root

	i := 1

	for i < len(line)-1 {
		if string(line[i]) == "[" {
			child := initNode()
			child.height = curr.height + 1
			child.parent = curr

			if curr.left == nil {
				curr.left = &child
				curr = curr.left
			} else if curr.right == nil {
				curr.right = &child
				curr = curr.right
			} else {
				panic(fmt.Sprintf("trying to insert 3 child nodes at i = %d", i))
			}
			i++
		} else if string(line[i]) == "]" || string(line[i]) == "," {
			curr = curr.parent
			i++
		} else {
			num := []byte{}
			for string(line[i]) != "," && string(line[i]) != "]" {
				num = append(num, line[i])
				i++
			}
			val := utils.ToInt(string(num))
			child := initNode()
			child.val = val
			child.height = curr.height + 1
			child.parent = curr

			if curr.left == nil {
				curr.left = &child
				curr = curr.left
			} else if curr.right == nil {
				curr.right = &child
				curr = curr.right
			} else {
				panic(fmt.Sprintf("trying to insert 3 child nodes at i = %d", i))
			}
		}
	}

	return &root
}
