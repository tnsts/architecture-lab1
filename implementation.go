package lab1

import (
		"fmt"
		"math"
		"strings"
		"strconv"
	)

func CalculatePostfix(input string) (int, error) {
	expression := strings.Fields(input)
	var stack []int

	for _, element := range expression {
		if element == "+" || element == "-" || element == "/" || element == "*" || element == "^" {
			if len(stack) < 2 {
				return 0, fmt.Errorf("Expression is wrong. Can't calculate: %s", element)
			}

			var values []int
			stack, values = stack[:len(stack) - 2], stack[len(stack) - 2:]
			a, b := values[0], values[1]

			switch element {
			case "+":
				stack = append(stack, a + b)
			case "-":
				stack = append(stack, a - b)
			case "*":
				stack = append(stack, a * b)
			case "/":
				stack = append(stack, a / b)
			case "^":
				stack = append(stack, int(math.Pow(float64(a), float64(b))))
			}
		} else {
			value, err := strconv.Atoi(element)
			if err != nil {
				return 0, fmt.Errorf("Invalid element of expression: %s", element)
			}

			stack = append(stack, value)
		}
	}

	if len(stack) != 1 {
		return 0, fmt.Errorf("Expression is not complete")
	}

	return stack[0], nil
}
