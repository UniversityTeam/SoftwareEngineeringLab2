package lab2

import (
	"fmt"
	"strings"
)

func ConvertPostfixToPrefix(inputString string) (string, error) {
	if (len(inputString) == 2 || len(inputString) == 1 || len(inputString) == 0) {
		return "", fmt.Errorf("Incorrect input postfix expression!!!")
	}
	var inputStringArray = strings.Split(inputString, " ")
	var operators = "+-*/"
	var stack []string

	for i := 0; i < len(inputStringArray); i++ {
		var currentElement = inputStringArray[i]
		if !strings.Contains(operators, currentElement) {
			stack = append(stack, currentElement)
		} else {
			var firstOperand = stack[len(stack)-1]
			stack = stack[:len(stack) - 1]
			var secondOperand = stack[len(stack)-1]
			stack = stack[:len(stack) - 1]

			var prefixForm = currentElement + " " + secondOperand + " " + firstOperand
			stack = append(stack, prefixForm)
		}

	}
	var result = strings.Join(stack, " ")
	return result, nil
}
