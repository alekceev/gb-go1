package main

import (
	"fmt"
	"math"
)

// 2. Задание для продвинутых (необязательное). Написать приложение, которое ищет все простые числа от 0 до N включительно.
// Число N должно быть задано из стандартного потока ввода.

func main() {
	var n int

	fmt.Print("Введите число: ")
	fmt.Scanln(&n)

	arr := init_prime_arr(n)

	fmt.Printf("Простые числа от 0 до %d: %v\n", n, arr)
}

func init_prime_arr(n int) []int {
	arr := []int{1}
	// для исключения чисел в степени
	exclude := map[int]struct{}{}
	for i := 2; i <= n; i++ {
		s := 1
		// заполняем числа в степени меньше n
		for {
			s++
			j := int(math.Pow(float64(i), float64(s)))
			if j > n {
				break
			}
			exclude[j] = struct{}{}
		}

		// числа в степени не попадают в результирующий массив
		if _, ok := exclude[i]; ok {
			continue
		}

		ok := true
		// исключаем числа, которые делятся без остатка на уже добавленные числа
		for _, j := range arr {

			if j == 1 {
				continue
			}

			if i%j == 0 {
				ok = false
				break
			}
		}

		if ok {
			arr = append(arr, i)
		}
	}
	return arr
}
