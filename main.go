package main

import (
	"fmt"
	"os"

	"github.com/17xande/bd-gator/internal/config"
)

type state struct {
	config *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		panic(err)
	}

	st := state{
		config: &cfg,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	rawArgs := os.Args
	if len(rawArgs) < 2 {
		fmt.Println("include a command")
		os.Exit(1)
	}

	commandName := rawArgs[1]
	args := rawArgs[2:]
	cmd := command{
		name: commandName,
		args: args,
	}
	err = cmds.run(&st, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
