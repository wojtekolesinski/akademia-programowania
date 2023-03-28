package main

import (
	"context"
	"fmt"
	"os"
	"reddit/fetcher"
	"sync"
	"time"
)

func main() {

	subreddits := []string{
		"golang",
		"python",
		"programming",
		"aww",
		"wtf",
		"iama",
		"bestof",
		"gaming",
		"pokemon",
		"minecraft",
		"skyrim",
	}

	headers := map[string]string{
		"User-agent": "golang-bot",
	}

	var wg sync.WaitGroup
	wg.Add(len(subreddits))

	for _, subreddit := range subreddits {
		go func(sub string) {
			defer wg.Done()
			var fetch fetcher.RedditFetcher
			fetch = &fetcher.HttpRedditFetcher{Url: fmt.Sprintf("http://reddit.com/r/%s.json", sub), Headers: headers}

			fmt.Printf("Fetching %s\n", sub)
			ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond * 5000)
			defer cancel()
			err := fetch.FetchWithContext(ctx)
			if err != nil {
				panic(err)
			}

			file, err := os.Create(fmt.Sprintf("./data/%s.txt", sub))
			if err != nil {
				panic(err)
			}

			fetch.Save(file)
			fmt.Printf("Saved %s\n", sub)
		}(subreddit)
	}
	wg.Wait()

	fmt.Println("DONE")
}
