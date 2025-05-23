package main

import (
	"context"
	"fmt"

	"github.com/17xande/bd-gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
		if err != nil {
			return fmt.Errorf("can't get user: %w", err)
		}

		return handler(s, cmd, user)
	}
}
