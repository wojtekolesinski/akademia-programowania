package main

import (
	"context"
	"fmt"
	"os"
	"reddit/fetcher"
	"sync"
	"time"
	"io"
)

func main() {
	var f fetcher.RedditFetcher // do not change
	var w io.Writer             // do not change

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
			f = &fetcher.HttpRedditFetcher{Url: fmt.Sprintf("http://reddit.com/r/%s.json", sub), Headers: headers}

			fmt.Printf("Fetching %s\n", sub)
			ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond * 5000)
			defer cancel()
			err := f.FetchWithContext(ctx)
			if err != nil {
				panic(err)
			}

			w, err := os.Create(fmt.Sprintf("./data/%s.txt", sub))
			if err != nil {
				panic(err)
			}

			f.Save(w)
			fmt.Printf("Saved %s\n", sub)
		}(subreddit)
	}
	wg.Wait()

	fmt.Println("DONE")
}
