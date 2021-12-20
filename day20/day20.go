package main

import (
	utils "aoc-2021/aoc-utils"
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

	algo := [512]int{}

	for i := 0; i < len(inputs[0]); i++ {
		if string(inputs[0][i]) == "." {
			algo[i] = 0
		} else {
			algo[i] = 1
		}
	}

	lines := strings.Split(inputs[1], "\r\n")

	padding := 2
	img := [104][104]int{}

	for j := 0; j < len(lines); j++ {

		for k := 0; k < len(lines[0]); k++ {
			if string(lines[j][k]) == "." {
				img[j+padding][k+padding] = 0
			} else {
				img[j+padding][k+padding] = 1
			}
		}
	}
	//printImg(img)
	newImg := applyAlg(img, algo[:])
	applyAlg(newImg, algo[:])

	//printImg(new2Img)
}

func getPixels(indexes [9][2]int, img [104][104]int) []int {
	pixels := []int{}
	for _, idx := range indexes {
		if idx[0] < 0 || idx[0] >= len(img) || idx[1] < 0 || idx[1] >= len(img[0]) {
			pixels = append(pixels, 0)
		} else {
			pixels = append(pixels, img[idx[0]][idx[1]])
		}
	}

	return pixels
}

func printImg(img [104][104]int) {
	for i := 0; i < len(img); i++ {
		for j := 0; j < len(img[0]); j++ {
			if img[i][j] == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func applyAlg(img [104][104]int, algo []int) [104][104]int {
	litMap := map[string]bool{}
	newImg := [104][104]int{}
	for i := 0; i < len(img); i++ {
		for j := 0; j < len(img[0]); j++ {

			indexes := [9][2]int{{i - 1, j - 1}, {i - 1, j}, {i - 1, j + 1}, {i, j - 1}, {i, j}, {i, j + 1}, {i + 1, j - 1}, {i + 1, j}, {i + 1, j + 1}}
			bin := getPixels(indexes, img)
			algoI := binToDecInt(bin)
			newImg[i][j] = algo[algoI]
			if algo[algoI] == 1 {
				key := utils.PointKey([]int{i, j})
				litMap[key] = true
			}
		}
	}
	fmt.Println(len(litMap))
	return newImg
}

//convert binary array to decimal int
func binToDecInt(binary []int) int {
	var dec int = 0
	for i := len(binary) - 1; i >= 0; i-- {
		digit := binary[i]
		pos := len(binary) - i - 1

		dec += digit * int(math.Pow(2, float64(pos)))
	}

	return dec
}
