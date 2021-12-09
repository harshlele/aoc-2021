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
	//fmt.Println(calcP1(lines))
	calcP2(lines)

}

func calcP1(lines []string) int {
	sum := 0

	for _, line := range lines {
		output := strings.Split(line, "|")
		words := strings.Split(output[1], " ")

		lMap := map[int]int{2: 1, 4: 4, 3: 7, 7: 8}

		for _, word := range words {
			_, ok := lMap[len(word)]
			if ok {
				sum += 1
			}
		}
	}

	return sum
}

func calcP2(lines []string) {
	wordArr := [][]string{}

	for _, line := range lines {
		output := strings.Replace(line, " | ", " ", -1)
		words := strings.Split(output, " ")
		wordArr = append(wordArr, words)
	}

	overall := 0

	for _, wArr := range wordArr {
		wMap := mapLetters(wArr)

		sum := 0
		for i := 10; i < 14; i++ {
			sum = (sum * 10) + getNumFromMap(wMap, wArr[i])
		}
		overall += sum
	}

	fmt.Println(overall)

}

/*
array - sequence mapping
		0
    ---------
1	|		|	2
	|	3	|
	---------
4	|		|	5
	|		|
	---------
		6
*/

func getInitSegMap() [7]map[string]bool {
	seg := [7]map[string]bool{}

	for i := 0; i < 7; i++ {
		seg[i] = map[string]bool{"a": true, "b": true, "c": true, "d": true, "e": true, "f": true, "g": true}
	}

	return seg
}

func mapLetters(words []string) [7]string {
	initSeg := getInitSegMap()

	completedMap := [7]string{"", "", "", "", "", "", ""}

	//first pass
	for _, w := range words {
		letters := strings.Split(w, "")

		offSeg := []int{}
		onSeg := []int{}

		//eliminate easy cases where you get the number from the length itself
		if len(w) == 2 {
			// displayed 1
			offSeg = []int{0, 1, 3, 4, 6}
			onSeg = []int{2, 5}
		} else if len(w) == 4 {
			//displayed 4
			offSeg = []int{0, 4, 6}
			onSeg = []int{1, 2, 3, 5}
		} else if len(w) == 3 {
			// displayed 7
			offSeg = []int{1, 3, 4, 6}
			onSeg = []int{0, 2, 5}
		} else if len(w) == 5 {
			//displayed 2,3 or 5
			onSeg = []int{0, 3, 6}
		} else if len(w) == 6 {
			//displayed 0,6 or 9
			onSeg = []int{0, 1, 5, 6}
		}

		//remove all the letters in the word from the off segment maps (eg if word is a 1, and contains a, then a can't be in segments 0,1,3,4 or 6)
		for _, l := range letters {
			for _, dInd := range offSeg {
				delete(initSeg[dInd], l)
			}

		}

		//remove all the letters in the on segment maps that are not in the word (ie if word is a 1 and contains b, then b can only be in segments 2 or 5, so remove b from all other segments)
		for _, dInd := range onSeg {
			for key := range initSeg[dInd] {
				if !strings.Contains(w, key) {
					delete(initSeg[dInd], key)
				}
			}
		}

		//if a map has only a single key, that means the letter has been mapped out, so store that
		for i, segment := range initSeg {
			if len(segment) == 1 {
				for k := range segment {
					completedMap[i] = k
				}
			}
		}

	}

	//use the completed maps(ie those with length 1) to remove duplicates from incomplete ones
	//(ie if a letter in an incompleted map has been completed previously, remove it from the incomplete map)
	for i, segment := range initSeg {
		if len(segment) > 1 {
			for _, m := range completedMap {
				if m != "" {
					_, ok := segment[m]
					if ok {
						delete(segment, m)
					}
				}
			}

			for k := range segment {
				completedMap[i] = k
			}
		}
	}

	return completedMap
}

//get number from word using the map
func getNumFromMap(lMap [7]string, word string) int {
	if len(word) == 2 {
		return 1
	} else if len(word) == 4 {
		return 4
	} else if len(word) == 3 {
		return 7
	} else if len(word) == 7 {
		return 8
	} else if len(word) == 5 {
		if !strings.Contains(word, lMap[2]) {
			return 5
		} else if strings.Contains(word, lMap[4]) {
			return 2
		} else {
			return 3
		}
	} else {
		if !strings.Contains(word, lMap[3]) {
			return 0
		} else if strings.Contains(word, lMap[4]) {
			return 6
		} else {
			return 9
		}
	}
}
