package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println(("Usage: calculator <number1> <operator> <number>"))
		return
	}

	num1, err1 := strconv.ParseFloat((os.Args[1]), 64)
	operator := os.Args[2]
	num2, err2 := strconv.ParseFloat((os.Args[3]), 64)

	if err1 != nil {
		fmt.Println("Please enter a valid number. Your input for the first number was not valid: ", os.Args[1])
		return
	}

	if err2 != nil {
		fmt.Println("Please enter a valid number. Your input for the second number was not valid: ", os.Args[2])
		return
	}

	result, calculationError := calculate(num1, operator, num2)

	if calculationError != nil {
		fmt.Println("An error occurred during calculation: ", calculationError)
		return
	}

	fmt.Println("The result of the calculation is: ", result)
}

func calculate(num1 float64, operator string, num2 float64) (float64, error) {
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "x":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("Cannot divide by zero")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("Invalid operator: %s", operator)
	}
}
