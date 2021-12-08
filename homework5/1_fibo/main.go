package main

import "fmt"

// 1. Напишите приложение, рекурсивно вычисляющее заданное из стандартного ввода число Фибоначчи.

func main() {
	var n int

	fmt.Print("Введите число: ")
	fmt.Scanln(&n)

	fmt.Printf("%d число ряда фибоначи = %d\n", n, fibo(n))
}

func fibo(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return fibo(n-1) + fibo(n-2)
}
