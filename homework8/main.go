package main

import (
	"flag"
	"fmt"
	"gb-go1/homework8/app/config"
	"log"
)

var configFile string

func main() {
	flag.StringVar(&configFile, "config", "", "path to config")
	flag.Parse()

	conf, err := config.Get(configFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(conf)
}
