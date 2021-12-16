package main

import (
	utils "aoc-2021/aoc-utils"
	"fmt"
	"io/ioutil"
	"math"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	hex := string(content)

	inp := hexToBin(hex)
	fmt.Println(scanBITSMsg(inp, 0))
}

func scanBITSMsg(binary []string, start int64) (int64, int64) {
	//id bits
	id := binStrToDecInt(binary[start+3 : start+6])
	//data packet
	if id == 4 {
		//first 6 bits for packet version and id, data starts from 7th bit
		upto := start + 6
		numCont := []string{}
		//if the bit is 0, that means this is the last number
		for binary[upto] != "0" {
			numCont = append(numCont, binary[upto+1:upto+5]...)
			upto += 5
		}
		//add the last digit and convert to decimal
		numCont = append(numCont, binary[upto+1:upto+5]...)
		val := binStrToDecInt(numCont)
		upto += 4

		return val, upto
	} else {
		//operation packet
		res := []int64{}
		var fUpto int64 = 0
		if binary[start+6] == "0" {
			//if the type bit is 0, we get the length of the contents of the packet in the next 15 bits
			l := binStrToDecInt(binary[start+7 : start+7+15])
			//run recursive scan of contents
			upto := start + 7 + 15
			curr := upto
			for upto < curr+l {
				a, b := scanBITSMsg(binary, upto)
				res = append(res, a)
				upto = b + 1
			}
			// subtract 1 because the last update for upto adds 1 to process the next packet, but there is no next packet,
			// so the last has to be subtracted away
			fUpto = upto - 1
		} else if binary[start+6] == "1" {
			//if the type bit is 1, we get the number of packets inside this packet in the next 11 bits
			l := binStrToDecInt(binary[start+7 : start+7+11])
			//loop until you have scanned l more packets
			var i int64 = 0
			upto := start + 7 + 11
			for i < l {
				a, b := scanBITSMsg(binary, upto)
				res = append(res, a)
				upto = b + 1
				i++
			}
			//same logic as above
			fUpto = upto - 1
		}
		// run the operation(based on id) on the values of the contents
		return runOp(res, id), fUpto
	}
}

//run operations on the data of the contents based on id
func runOp(data []int64, id int64) int64 {
	if id == 0 {
		var sum int64 = 0
		for _, num := range data {
			sum += num
		}
		return sum
	} else if id == 1 {
		var prod int64 = 1
		for _, num := range data {
			prod *= num
		}
		return prod
	} else if id == 2 {
		var min int64 = math.MaxInt
		for _, num := range data {
			if num < min {
				min = num
			}
		}
		return min
	} else if id == 3 {
		var max int64 = math.MinInt
		for _, num := range data {
			if num > max {
				max = num
			}
		}
		return max
	} else if id == 5 {
		if data[0] > data[1] {
			return 1
		} else {
			return 0
		}
	} else if id == 6 {
		if data[0] < data[1] {
			return 1
		} else {
			return 0
		}
	} else {
		if data[0] == data[1] {
			return 1
		} else {
			return 0
		}
	}
}

//convert hex string to binary array
func hexToBin(hex string) []string {
	binMap := map[string][4]string{
		"0": {"0", "0", "0", "0"},
		"1": {"0", "0", "0", "1"},
		"2": {"0", "0", "1", "0"},
		"3": {"0", "0", "1", "1"},
		"4": {"0", "1", "0", "0"},
		"5": {"0", "1", "0", "1"},
		"6": {"0", "1", "1", "0"},
		"7": {"0", "1", "1", "1"},
		"8": {"1", "0", "0", "0"},
		"9": {"1", "0", "0", "1"},
		"A": {"1", "0", "1", "0"},
		"B": {"1", "0", "1", "1"},
		"C": {"1", "1", "0", "0"},
		"D": {"1", "1", "0", "1"},
		"E": {"1", "1", "1", "0"},
		"F": {"1", "1", "1", "1"},
	}

	binary := make([]string, len(hex)*4)

	for i := 0; i < len(hex); i++ {
		b := binMap[string(hex[i])]
		binary[4*i], binary[4*i+1], binary[4*i+2], binary[4*i+3] = b[0], b[1], b[2], b[3]
	}

	return binary
}

//convert binary array to decimal int
func binStrToDecInt(binary []string) int64 {
	var dec int64 = 0
	for i := len(binary) - 1; i >= 0; i-- {
		digit := int64(utils.ToInt(string(binary[i])))
		pos := float64(len(binary) - i - 1)

		dec += digit * int64(math.Pow(2, pos))
	}

	return dec
}
