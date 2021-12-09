package main

import (
	"errors"
	"fmt"
	"io"
)

// 1. С помощью структур и указателей реализуйте структру данных линейный связный список,
// заполните его с потока ввода (как в задании с сортировкой массива).
// Затем реализуйте алгоритм разворота односвязного списка.

type Node struct {
	val  int
	next *Node
}

type List struct {
	head *Node
}

func CreateList() *List {
	return &List{head: &Node{}}
}

func (l *List) GetLast() *Node {
	n := l.head
	for n != nil && n.next != nil {
		n = n.next
	}
	return n
}

func (l *List) Push(val int) {
	n := l.GetLast()
	n.next = &Node{
		val: val,
	}
}

func (l *List) String() string {
	n := l.head.next
	var s string
	for n != nil {
		s += fmt.Sprintf("%v ", n.val)
		n = n.next
	}
	return s
}

func (l *List) Reverse() {
	n := l.head.next
	var prev *Node
	var next *Node
	for n != nil {
		next = n.next
		n.next = prev
		prev = n
		n = next
	}
	l.head.next = prev
}

func main() {
	list := CreateList()

	fmt.Println("Введите числа, затем нажмите ctrl+D")

	for {
		var i int
		if _, err := fmt.Scanln(&i); errors.Is(err, io.EOF) {
			fmt.Println("Exit")
			break
		}

		list.Push(i)
	}

	fmt.Println(list)
	list.Reverse()
	fmt.Println(list)
}
