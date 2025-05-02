package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/17xande/bd-gator/internal/config"
)

type state struct {
	config *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	list map[string]func(*state, command) error
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("supply username as argument to this command")
	}

	if err := s.config.SetUser(cmd.args[0]); err != nil {
		return fmt.Errorf("error saving user to config file: %w", err)
	}

	fmt.Println("User has been set.")
	return nil
}

func getCommands() commands {
	return commands{
		list: make(map[string]func(*state, command) error),
	}
}

func (c *commands) run(s *state, cmd command) error {
	command, exists := c.list[cmd.name]
	if !exists {
		return errors.New("command doesn't exist")
	}

	return command(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.list[name] = f
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		panic(err)
	}

	st := state{
		config: &cfg,
	}

	commands := getCommands()
	commands.register("login", handlerLogin)

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
	err = commands.run(&st, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
