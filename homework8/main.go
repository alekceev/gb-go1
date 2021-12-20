package main

import (
	"fmt"
	"gb-go1/homework8/app/config"
	"log"
)

func main() {
	conf, err := config.Get()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(conf)
}
