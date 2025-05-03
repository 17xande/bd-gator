package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/17xande/bd-gator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("supply username as argument to this command")
	}

	_, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("error checking if user exists: %w", err)
	}

	if err := s.config.SetUser(cmd.args[0]); err != nil {
		return fmt.Errorf("error saving user to config file: %w", err)
	}

	fmt.Println("User has been set.")
	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error getting all users: %w", err)

	}

	for _, u := range users {
		current := ""
		if s.config.CurrentUserName == u.Name {
			current = " (current)"
		}
		fmt.Printf(" * %s%s\n", u.Name, current)
	}

	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("supply username as argument to this command")
	}

	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	}
	user, err := s.db.CreateUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error creating user in database: %w", err)
	}

	fmt.Printf("user created\n%v\n", user)

	if err := s.config.SetUser(cmd.args[0]); err != nil {
		return fmt.Errorf("error saving user to config file: %w", err)
	}

	return nil
}

func handlerReset(s *state, cmd command) error {
	return s.db.Reset(context.Background())
}
