package main

import (
	"errors"
	"fmt"
	"io"
)

// 1. Познакомьтесь с алгоритмом сортировки вставками.
// Напишите приложение, которое принимает на вход набор целых чисел и отдаёт на выходе его же в отсортированном виде.
// Сохраните код, он понадобится нам в дальнейших уроках.

func main() {
	arr := []int{}
	fmt.Println("Введите числа, затем нажмите ctrl+D")
	for {
		var i int
		if _, err := fmt.Scanln(&i); errors.Is(err, io.EOF) {
			fmt.Println("Exit")
			break
		}
		arr = append(arr, i)
	}

	// sort
	for i := 0; i <= len(arr)-1; i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}

	fmt.Println(arr)
}
