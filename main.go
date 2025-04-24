package main

import (
	"fmt"

	"github.com/17xande/bd-gator/internal/config"
)

func main() {
	cfg := config.Read()
	cfg.SetUser("lane")
	cfg = config.Read()
	fmt.Printf("%v\n", cfg)
}
