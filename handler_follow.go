package main

import (
	"context"
	"fmt"
	"time"

	"github.com/17xande/bd-gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.name)
	}

	// get feed id based on feed url (1st arg)
	feed, err := s.db.GetFeed(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("can't get feed from url: %w", err)
	}

	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}
	follow, err := s.db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return fmt.Errorf("can't create feed follow: %w", err)
	}

	fmt.Printf("%v\n", follow)

	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("usage: %s", cmd.name)
	}

	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("can't get feeds for user: %w", err)
	}

	fmt.Printf("Feeds for %s:\n%v\n", s.config.CurrentUserName, feeds)

	return nil
}
