package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func inputNumFile(fileName string, valArr *[]float64) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	strs := strings.Fields(string(file))

	if len(strs) != 3 {
		fmt.Print("Error: Invalid file format, expected 3 values\n")
		os.Exit(1)
	}

	for indx, val := range strs {
		num, err := strconv.ParseFloat(val, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}

		if indx == 0 && num == 0 {
			fmt.Print("Error. a cannot be 0\n")
			os.Exit(1)
		}

		if err == nil {
			*valArr = append(*valArr, num)
		}
	}
}

func inputNumConsole(valArr *[]float64) {
	var num string
	nameArr := []string{"A", "B", "C"}

	for i := 0; i < len(nameArr); i++ {
		fmt.Printf("%s: ", nameArr[i])
		fmt.Fscan(os.Stdin, &num)
		input, err := strconv.ParseFloat(num, 64)
		if err != nil || i == 0 && input == 0 {
			fmt.Print("Error. Expected a valid real number, got invalid instead")
			os.Exit(0)
		}
		*valArr = append(*valArr, input)
	}
}

func printAnswer(valArr, answerMas []float64, discrim float64) {
	fmt.Printf("Equation is: (%g) x^2 + (%g) x + (%g) = 0 \n", valArr[0], valArr[1], valArr[2])
	fmt.Printf("There are %d roots \n", len(answerMas))

	for idx, val := range answerMas {
		fmt.Printf("x%d = %.2f\n", idx+1, val)
	}
}

func calculationResult(valArr []float64, discrim *float64) []float64 {
	var answerMas []float64

	if *discrim = (valArr[1]*valArr[1] - 4*valArr[0]*valArr[2]); *discrim > 0 {
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
	var fileName string

	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}

	if fileName != "" {
		inputNumFile(fileName, &valArr)
	} else {
		inputNumConsole(&valArr)
	}

	answerMas := calculationResult(valArr, &discrim)
	printAnswer(valArr, answerMas, discrim)
}
