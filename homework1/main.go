package main

import "fmt"

type User struct {
	Name string
}

func (u User) String() string {
	return u.Name
}

func main() {
	u := User{
		Name: "World",
	}

	fmt.Println("Hello", u)
}
