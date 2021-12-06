package main

import "fmt"

// 1. Напишите приложение, рекурсивно вычисляющее заданное из стандартного ввода число Фибоначчи.

func main() {
	var n int

	fmt.Print("Введите число: ")
	fmt.Scanln(&n)

	fibo := func() func() int {
		first, second := 0, 1
		return func() int {
			defer func() { first, second = second, first+second }()
			return first
		}
	}

	f := fibo()

	for i := 0; i <= n; i++ {
		fmt.Println(f())
	}
}
