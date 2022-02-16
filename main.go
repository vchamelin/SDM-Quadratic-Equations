package main

import "fmt"

func main() {
	var numA, numB uint

	fmt.Print("Введите число A: ")
	fmt.Scan(&numA)
	fmt.Print("Введите число B: ")
	fmt.Scan(&numB)

	fmt.Printf("Вы ввели числа %d и %d", numA, numB)
}
