package main

import (
	"fmt"
)

// 3. С клавиатуры вводится трехзначное число. Выведите цифры, соответствующие количество сотен, десятков и единиц в этом числе.

func main() {
	var n int

	fmt.Println("Введите число:")
	fmt.Scanln(&n)

	names := map[int]string{
		0: "Едениц",
		1: "Десятков",
		2: "Сотен",
		3: "Тысяч",
	}

	var i int
	for {
		if n == 0 {
			break
		}
		s := n % 10
		n = n / 10

		name := names[i]
		if name == "" {
			name = fmt.Sprintf("%dй разряд", i)
		}

		fmt.Printf("%s: %d\n", name, s)
		i++
	}

}
