package main

import (
	"fmt"

	"github.com/17xande/bd-gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		panic(err)
	}

	err = cfg.SetUser("lane")
	if err != nil {
		panic(err)
	}
	cfg, err = config.Read()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", cfg)
}
