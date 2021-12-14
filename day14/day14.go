package main

import (
	"fmt"
	"io/ioutil"
	"math"
	re "regexp"
	"strings"
)

func main() {

	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	exp := re.MustCompile("\\r\\n\\r\\n")

	inputs := exp.Split(string(content), -1)
	template := inputs[0]
	lines := strings.Split(inputs[1], "\r\n")
	rules := map[string]string{}
	//make a rules map
	for _, l := range lines {
		k := strings.Split(l, " -> ")
		rules[k[0]] = k[1]
	}

	//initMap contains the number of occurences of letter pairs
	//in the original template string (NN, NC, CB in the test input)
	initMap := map[string]int64{}
	countMap := map[string]int64{}
	for i := 0; i < len(template); i++ {
		if i < len(template)-1 {
			//if a pair is in the rules, store its count in initMap
			_, ok := rules[template[i:i+2]]
			if ok {
				_, o2 := initMap[template[i:i+2]]
				if !o2 {
					initMap[template[i:i+2]] = 1
				} else {
					initMap[template[i:i+2]] += 1
				}

			}
		}

		_, o2 := countMap[string(template[i])]
		if !o2 {
			countMap[string(template[i])] = 1
		} else {
			countMap[string(template[i])] += 1
		}
	}

	calcPoly(initMap, rules, countMap, 40)

}

func calcPoly(initMap map[string]int64, rules map[string]string, counts map[string]int64, steps int) {

	for i := 0; i < steps; i++ {
		//new map for every step
		newMap := map[string]int64{}
		for k := range initMap {
			//check if pair is in initMap
			v, o1 := rules[k]
			if o1 {
				//get the letter to insert, and store the occurences of the 2 new pairs that will be generated
				//eg NN -> C (ie NCN after inserting C)
				//so store the count of NC and CN in newMap
				l := strings.Split(k, "")
				k1, k2 := l[0]+v, v+l[1]

				_, o2 := newMap[k1]
				if !o2 {
					newMap[k1] = initMap[k]
				} else {
					newMap[k1] += initMap[k]
				}

				_, o3 := newMap[k2]
				if !o3 {
					newMap[k2] = initMap[k]
				} else {
					newMap[k2] += initMap[k]
				}

				//update counts
				_, o4 := counts[v]
				if !o4 {
					counts[v] = initMap[k]
				} else {
					counts[v] += initMap[k]
				}
			}
		}
		//replace old map with new one
		initMap = newMap
	}

	//get the max and min used letters, and calc answer
	var max, min int64
	max = 0
	min = math.MaxInt64

	for k := range counts {
		if max < counts[k] {
			max = counts[k]
		}
		if min > counts[k] {
			min = counts[k]
		}
	}

	fmt.Println("diff using counts", max-min)
}
