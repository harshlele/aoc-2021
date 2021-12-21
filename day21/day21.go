package main

import "fmt"

func main() {
	p1, p1Score := 4, 0
	p2, p2Score := 6, 0

	moves := 0

	for true {
		shift1 := detDie(moves)
		p1 = (p1-1+shift1)%10 + 1
		p1Score += p1
		moves++

		if p1Score >= 1000 {
			break
		}

		shift2 := detDie(moves)
		p2 = (p2-1+shift2)%10 + 1
		p2Score += p2
		moves++

		if p2Score >= 1000 {
			break
		}
	}
	fmt.Println("P1 score", p1Score)
	fmt.Println("P2 score", p2Score)
}

func detDie(m int) int {
	i1, i2, i3 := (3*m)%100, (3*m+1)%100, (3*m+2)%100
	return i1 + i2 + i3 + 3
}
