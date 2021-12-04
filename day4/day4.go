package main

import (
	"fmt"
	"io/ioutil"
	re "regexp"
	"strconv"
	"strings"
)

func main() {
	draws := []int{}
	cards := [][]int{}

	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	reg := re.MustCompile("\\r\\n\\r\\n")
	inputs := reg.Split(string(content), -1)

	draws = parseDraws(inputs[0])
	cards = parseCards(inputs[1:])

	bingoCard := []int{}
	bingoDraw := 0

	for i := 0; i < len(draws); i++ {
		done := false
		for j := 0; j < len(cards); j++ {
			index := markNumber(cards[j], draws[i])
			if index == -1 {
				continue
			}
			if checkIfBingo(cards[j], index) {
				bingoCard = cards[j]
				bingoDraw = draws[i]
				done = true
				break
			}
		}
		if done {
			break
		}

	}

	fmt.Println(calcFirstPart(bingoCard, bingoDraw))

}

func parseDraws(line string) []int {
	dArr := []int{}
	nos := strings.Split(line, ",")
	for i := 0; i < len(nos); i++ {
		a, err := strconv.ParseInt(nos[i], 10, 0)
		if err != nil {
			panic(err)
		}
		dArr = append(dArr, int(a))
	}
	return dArr
}

func parseCards(iArr []string) [][]int {
	cards := [][]int{}

	for i := 0; i < len(iArr); i++ {
		cardStr := strings.Fields(iArr[i])
		card := []int{}
		for j := 0; j < len(cardStr); j++ {
			a, err := strconv.ParseInt(cardStr[j], 10, 0)
			if err != nil {
				panic(err)
			}
			card = append(card, int(a))
		}
		cards = append(cards, card)
	}

	return cards
}

func calcFirstPart(card []int, draw int) int {
	sum := 0
	for i := 0; i < len(card); i++ {
		if card[i] != -1 {
			sum += card[i]
		}
	}
	return sum * draw
}

//for unused variables lol
func x(X ...interface{}) {}

func markNumber(card []int, no int) int {

	for i := 0; i < len(card); i++ {
		if card[i] == no {
			card[i] = -1
			return i
		}
	}
	return -1
}

func checkIfBingo(card []int, i int) bool {

	mult := (i / 5) * 5

	row := []int{card[i], card[(i+1)%5+mult], card[(i+2)%5+mult], card[(i+3)%5+mult], card[(i+4)%5+mult]}
	col := []int{card[i], card[(i+5)%25], card[(i+10)%25], card[(i+15)%25], card[(i+20)%25]}

	bingo := checkSlice(row) || checkSlice(col)

	if bingo {
		return true
	}

	return false
}

func checkSlice(slice []int) bool {

	for i := 0; i < len(slice); i++ {
		if slice[i] != -1 {
			return false
		}
	}
	return true
}
