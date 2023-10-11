package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/ntvviktor/BlogApplication/internal/database"
)

func startScraping(db *database.Queries, concurrency int, timeBetweenRequest time.Duration) {
	log.Printf("Start scarping data on %v for each %s \n", concurrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFetchToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Printf("Error occurs when scraping")
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(wg, db, feed)
		}
		wg.Wait()
	}

}

func scrapeFeed(wg *sync.WaitGroup, db *database.Queries, feed database.Feed) {
	defer wg.Done()

	_, err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("Error when marking feed as fetched")
		return
	}
	rssFeed, err := URLToFeed(feed.Url)
	if err != nil {
		log.Fatal("Error when fetching data")
	}
	for _, item := range rssFeed.Channel.Item {
		log.Println(item.Title)
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}
