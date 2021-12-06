package main

import "fmt"

// 2. Напишите приложение, рекурсивно вычисляющее заданное из стандартного ввода число Фибоначчи.
// Оптимизируйте приложение за счёт сохранения предыдущих результатов в мапе.

func main() {
	var n int

	fmt.Print("Введите число: ")
	fmt.Scanln(&n)

	fibo := map[string]int{"first": 0, "second": 1}

	f := func() int {
		defer func() { fibo["first"], fibo["second"] = fibo["second"], fibo["first"]+fibo["second"] }()
		return fibo["first"]
	}

	for i := 0; i <= n; i++ {
		fmt.Println(f())
	}
}
