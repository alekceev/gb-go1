package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

// 1. Доработать калькулятор: больше операций и валидации данных (проверка деления на 0, возведение в степень, факториал)
//  + форматирование строки при выводе дробного числа ( 2 знака после точки)

func main() {
	var a, b, res float64
	var op string

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)

	ops := []string{"+", "-", "*", "/", "**", "!"}

	fmt.Print("Введите арифметическую операцию " + strings.Join(ops, ", ") + ": ")
	fmt.Scanln(&op)

	if op != "!" {
		fmt.Print("Введите второе число: ")
		fmt.Scanln(&b)
	}

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if b == 0 {
			fmt.Println("Делить на ноль нельзя")
			os.Exit(1)
		}
		res = a / b
	case "**":
		res = math.Pow(a, b)
	case "!":
		res = factorial(a)
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}

	fmt.Printf("Результат выполнения операции: %.2f\n", res)
}

func factorial(n float64) float64 {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
