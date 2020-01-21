package main

import (
	"strings"
	"fmt"
	"strconv"
)


func main(){
	inputString := "22+30"
	input := strings.Split(inputString, "+")
	fmt.Println("Final Input: ", input)
	op1, _ := parseString(input[0])
	op2, err := strconv.Atoi(input[1])
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println("Result: ", op1 + op2)
	operate("+", 5, 10)
	operate("*", 2, 8)
	operate("-", 20, 30)
	operate("p", 10, 25)
}

func parseString(operator string) (int, error) {
	result, err := strconv.Atoi(operator)
	return result, err					
}

func operate(operation string, value1 int, value2 int) {
	switch operation{
		case "+":
			fmt.Println("Add: ", value1 + value2)
		case "-":
			fmt.Println("Sub: ", value1 - value2)
		case "*":
			fmt.Println("Mult: ", value1 * value2)
		case "/":
			fmt.Println("Div: ", value1 / value2)
		default:
			fmt.Println(operation, "operation is not supported!")
			
	}

}