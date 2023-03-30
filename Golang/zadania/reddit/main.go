package main

import (
	"context"
	"fmt"
	"golang.org/x/exp/slog"
	"io"
	"os"
	"reddit/fetcher"
	"sync"
	"time"
)

func main() {
	//var f fetcher.RedditFetcher // do not change
	var w io.Writer // do not change

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
		"asdfssdaa",
	}

	headers := map[string]string{
		"User-agent": "golang-bot",
	}

	var wg sync.WaitGroup
	wg.Add(len(subreddits))

	for _, subreddit := range subreddits {
		go func(sub string) {
			defer wg.Done()
			f := &fetcher.HttpRedditFetcher{Url: fmt.Sprintf("https://reddit.com/r/%s.json", sub), Headers: headers}

			slog.Info(fmt.Sprintf("Fetching %s\n", sub))
			ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*1000)
			defer cancel()
			err := f.FetchWithContext(ctx)
			if err != nil {
				slog.Error(err.Error())
				return
			}

			w, err = os.Create(fmt.Sprintf("./data/%s.txt", sub))
			if err != nil {
				slog.Error(err.Error())
				return
			}

			err = f.Save(w)
			if err != nil {
				slog.Error(err.Error())
				return
			}
			slog.Info(fmt.Sprintf("Saved %s\n", sub))
		}(subreddit)
	}
	wg.Wait()

	slog.Info("DONE")
}
