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
	completedCards := map[int]bool{}
	p2 := true

	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	//split lines
	reg := re.MustCompile("\\r\\n\\r\\n")
	inputs := reg.Split(string(content), -1)
	//get draws and cards arrays
	draws = parseDraws(inputs[0])
	cards = parseCards(inputs[1:])

	bingoCard := []int{}
	bingoDraw := 0

	for i := 0; i < len(draws); i++ {
		done := false
		for j := 0; j < len(cards); j++ {
			//if solving part 2 and the cards has already been completed, skip the card
			if p2 {
				_, ok := completedCards[j]
				if ok {
					continue
				}

			}
			//mark the drawn number, if no num was marked skip the card
			index := markNumber(cards[j], draws[i])
			if index == -1 {
				continue
			}

			//if its p1 and bingo, just have to check if its the first bingo, so store card and break
			//if p2, mark the card as completed in the map
			//if all cards have been completed, current card is the last one, so store that and break
			if checkIfBingo(cards[j], index) {
				bingoDraw = draws[i]

				if p2 == true {
					completedCards[j] = true
					if len(completedCards) == len(cards) {
						bingoCard = cards[j]
						done = true
						break
					}
				} else {
					bingoCard = cards[j]
					done = true
					break
				}
			}
		}
		if done {
			break
		}

	}

	fmt.Println(calcResult(bingoCard, bingoDraw))

}

//read draws into array
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

//read cards
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

func calcResult(card []int, draw int) int {
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

// check if a card has bingo'd
func checkIfBingo(card []int, i int) bool {
	//row offset
	mult := (i / 5) * 5
	//get all numbers in the same row(%5 to wrap around)
	row := []int{card[i], card[(i+1)%5+mult], card[(i+2)%5+mult], card[(i+3)%5+mult], card[(i+4)%5+mult]}
	//get all numbers in the same column
	col := []int{card[i], card[(i+5)%25], card[(i+10)%25], card[(i+15)%25], card[(i+20)%25]}

	return checkSlice(row) || checkSlice(col)

}

func checkSlice(slice []int) bool {

	for i := 0; i < len(slice); i++ {
		if slice[i] != -1 {
			return false
		}
	}
	return true
}
