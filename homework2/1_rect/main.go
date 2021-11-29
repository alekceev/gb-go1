package main

import "fmt"

// 1. Напишите программу для вычисления площади прямоугольника. Длины сторон прямоугольника должны вводиться пользователем с клавиатуры.

func main() {
	var a, b float64

	fmt.Println("Чтобы вычислить площадь прямоугольник введите")
	fmt.Println("сторону A")
	fmt.Scanln(&a)
	fmt.Println("сторону B")
	fmt.Scanln(&b)
	fmt.Println("Площадь равна:", a*b)
}
