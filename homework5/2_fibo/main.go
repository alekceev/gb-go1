package main

import "fmt"

// 2. Напишите приложение, рекурсивно вычисляющее заданное из стандартного ввода число Фибоначчи.
// Оптимизируйте приложение за счёт сохранения предыдущих результатов в мапе.

func main() {
	var n int

	fmt.Print("Введите число: ")
	fmt.Scanln(&n)

	cache := map[string]int{"first": 0, "second": 1, "n": 1}
	fibo(n, cache)
	fmt.Printf("%d число ряда фибоначи = %d\n", n, cache["second"])
}

func fibo(n int, cache map[string]int) {
	if cache["n"] < n {
		cache["n"]++
		fibo(n, cache)
		cache["first"], cache["second"] = cache["second"], cache["first"]+cache["second"]
	}
}
