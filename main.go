package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func addNum(valArr *[]float64) bool {
	var num string
	nameArr := []string{"A", "B", "C"}

	for i := 0; i < len(nameArr); i++ {
		fmt.Printf("%s: ", nameArr[i])
		fmt.Fscan(os.Stdin, &num)
		input, err := strconv.ParseFloat(num, 64)
		if err != nil || i == 0 && input == 0 {
			fmt.Print("Error. Expected a valid real number, got invalid instead")
			return false
		}
		*valArr = append(*valArr, input)
	}
	return true
}

func printAnswer(valArr, answerMas []float64, discrim float64) {
	fmt.Printf("Equation is: (%g) x^2 + (%g) x + (%g) = %g \n", valArr[0], valArr[1], valArr[2], discrim)
	fmt.Printf("There are %d roots \n", len(answerMas))

	for idx, val := range answerMas {
		fmt.Printf("x%d = %.2f\n", idx+1, val)
	}
}

func calculationResult(valArr []float64, discrim *float64) []float64 {
	var answerMas []float64

	if *discrim = (valArr[1]*valArr[1] - 4*valArr[0]*valArr[2]); *discrim >= 0 {
		firstX := (-valArr[1] + math.Sqrt(*discrim)) / (2 * valArr[0])
		secondX := (-valArr[1] - math.Sqrt(*discrim)) / (2 * valArr[0])
		if firstX == secondX {
			answerMas = append(answerMas, firstX)
		} else {
			answerMas = append(answerMas, firstX, secondX)
		}
	} else if *discrim == 0 {
		firstX := -valArr[1] / (2 * valArr[0])
		answerMas = append(answerMas, firstX)
	}
	return answerMas
}

func main() {
	var discrim float64
	var valArr []float64

	if addNum(&valArr) {
		answerMas := calculationResult(valArr, &discrim)
		printAnswer(valArr, answerMas, discrim)
	}

}
