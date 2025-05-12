package main

import (
	"context"
	"fmt"
	"time"

	"github.com/17xande/bd-gator/internal/rss"
)

func handlerAggregator(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <time_between_reqs>", cmd.name)
	}

	dur, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("can't parse duration: %w", err)
	}

	fmt.Printf("collecting feeds every %s\n", cmd.args[0])

	ticker := time.NewTicker(dur)
	for ; ; <-ticker.C {
		scrapeFeeds(s, cmd)
	}

}

func scrapeFeeds(s *state, cmd command) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't fetch next feed: %w", err)
	}

	if err := s.db.MarkFeedFetched(context.Background(), feed.ID); err != nil {
		return fmt.Errorf("could't mark feed as fetched: %w", err)
	}

	rssfeed, err := rss.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}

	for _, item := range rssfeed.Channel.Item {
		fmt.Printf("%s\n", item.Title)
	}
	return nil
}
