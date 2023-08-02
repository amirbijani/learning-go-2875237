package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {

	float1 := getInput("Value 1: ")
	float2 := getInput("Value 2: ")

	var result float64
	switch operation := getOperation(); operation {
	case "+":
		result = Sum(float1, float2)
	case "-":
		result = Minus(float1, float2)
	case "/":
		result = Devide(float1, float2)
	case "*":
		result = Multiply(float1, float2)
	default:
		panic("Invalid operation")
	}
	result = math.Round(result*100) / 100
	fmt.Printf("The result is %v\n\n", result)

}

func getInput(prompt string) float64 {
	fmt.Printf("%v", prompt)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		message := fmt.Sprintf("%v must be a number", prompt)
		panic(message)
	}
	return (value)
}

func getOperation() string {
	fmt.Print("Seleect an operation (+ - * /) : ")
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}

func Sum(value1, value2 float64) float64 {

	return value1 + value2
}

func Minus(value1, value2 float64) float64 {

	return value1 - value2
}

func Devide(value1, value2 float64) float64 {
	return value1 / value2
}

func Multiply(value1, value2 float64) float64 {
	return value1 * value2
}
