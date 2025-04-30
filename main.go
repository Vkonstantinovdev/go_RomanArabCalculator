package main

import (
	"fmt"
)

func main() {
	result, err := checkOperand(10, 0, "/")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", result)
	}

	result, err = checkOperand(10, 2, "/")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", result)
	}
}

func add(operator1, operator2 float64) float64 {
	return operator1 + operator2
}

func smultiply(operator, operator2 float64) float64 {
	return operator * operator2
}

func subtract(operator, operator2 float64) float64 {
	return operator - operator2
}

func divide(operator, operator2 float64) (float64, error) {
	if operator2 == 0 {
		return 0, fmt.Errorf("деление на ноль при операнде %.0f", operator2)
	}
	return operator / operator2, nil
}

func checkOperand(operator, operator2 float64, operand string) (float64, error) {
	switch operand {
	case "+":
		return add(operator, operator2), nil
	case "-":
		return subtract(operator, operator2), nil
	case "*":
		return smultiply(operator, operator2), nil
	case "/":
		result, err := divide(operator, operator2)
		if err != nil {
			return 0, fmt.Errorf("ошибка при делении %.0f на %.0f: %w", operator, operator2, err)
		}
		return result, nil
	default:
		return 0, fmt.Errorf("неизвестный оператор: %s", operand)
	}
}

