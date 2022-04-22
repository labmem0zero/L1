package main

import (
	"fmt"
)

func SolveProblem(f func(), probNumb int){
	fmt.Printf("- - - - - - - - - Задание №\t%v: - - - - - - - - - - -\n",probNumb)
	f()
	fmt.Printf("- - - - - - - - - - - - - - - - - - - - - - - - - - -\n")
	fmt.Println()
}

func main() {
	SolveProblem(prob1, 1)
	SolveProblem(prob2, 2)
	SolveProblem(prob3,3)
	SolveProblem(prob6, 6)
	SolveProblem(prob7, 7)
	SolveProblem(prob8,8)
	SolveProblem(prob9,9)
	SolveProblem(prob10,10)
	SolveProblem(prob11, 11)
	SolveProblem(prob12, 12)
	SolveProblem(prob13, 13)
	SolveProblem(prob14,14)
	SolveProblem(prob15,15)
	SolveProblem(prob16,16)//ОСМЫСЛИТЬ
	SolveProblem(prob17,17)
	SolveProblem(prob18,18)
	SolveProblem(prob19,19)
	SolveProblem(prob20,20)
	SolveProblem(prob21,21)//ПЕРЕОСМЫСЛИТЬ
	SolveProblem(prob22,22)
	SolveProblem(prob23,23)
	SolveProblem(prob24,24)
	SolveProblem(prob25,25)
	SolveProblem(prob26, 26)

}