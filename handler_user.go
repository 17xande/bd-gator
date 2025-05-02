package main

import (
	"errors"
	"fmt"
)

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
