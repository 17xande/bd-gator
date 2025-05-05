package main

import (
	"context"
	"fmt"

	"github.com/17xande/bd-gator/internal/rss"
)

func handlerAggregator(s *state, cmd command) error {
	feed, err := rss.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}

	fmt.Printf("%v\n", feed)
	return nil
}
