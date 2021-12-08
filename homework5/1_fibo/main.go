package main

import "fmt"

// 1. Напишите приложение, рекурсивно вычисляющее заданное из стандартного ввода число Фибоначчи.

func main() {
	var n int

	fmt.Print("Введите число: ")
	fmt.Scanln(&n)

	arr := fibo([]int{}, n)
	fmt.Println(arr)
}

func fibo(arr []int, n int) []int {
	if len(arr) == 0 {
		arr = append(arr, 0, 1)
	}
	arr = append(arr, arr[len(arr)-2]+arr[len(arr)-1])
	if len(arr) < n {
		arr = fibo(arr, n)
	}
	return arr
}
