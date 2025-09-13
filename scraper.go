package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/jupitters/rssagg/internal/database"
)

func startScrapping(
	db *database.Queries,
	concurrency int,
	timeBetweenRequests time.Duration,
) {
	log.Printf("Scrapping on %v goroutines every %s duration", concurrency, timeBetweenRequests)
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(
			context.Background(),
			int32(concurrency),
		)
		if err != nil {
			log.Println("Erro buscando feeds: %v", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)

			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("Erro marcando o feed como fetched:", err)
		return
	}

	rssFeed, err := URLtoFeed(feed.Url)
	if err != nil {
		log.Println("Erro no fetching do feed:", err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		log.Println("Post encontrado:", item.Title)
	}
	log.Printf("Feed %s buscado, %v posts encontrados,", feed.Name, len(rssFeed.Channel.Item))

}
