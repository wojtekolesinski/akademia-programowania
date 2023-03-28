package fetcher

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type response struct {
	Data struct {
		Children []struct {
			Data struct {
				Title string `json:"title"`
				URL   string `json:"url"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

type RedditFetcher interface {
	Fetch() error
	FetchWithContext(ctx context.Context) error
	Save(io.Writer) error
}

type HttpRedditFetcher struct {
	Url     string
	Headers map[string]string
	resp    response
}

func (fetcher *HttpRedditFetcher) Fetch() error {
	req, err := http.NewRequest("GET", fetcher.Url, nil)
	if err != nil {
		return err
	}

	for key, value := range fetcher.Headers {
		req.Header.Add(key, value)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	(*fetcher).resp = response{}
	err = json.Unmarshal(body, &fetcher.resp)
	return err
}

func (fetcher *HttpRedditFetcher) FetchWithContext(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, "GET", fetcher.Url, nil)
	if err != nil {
		return err
	}
	
	for key, value := range fetcher.Headers {
		req.Header.Add(key, value)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	(*fetcher).resp = response{}
	err = json.Unmarshal(body, &fetcher.resp)
	return err
}

func (fetcher* HttpRedditFetcher) Save(writer io.Writer) error {
	for _, child := range fetcher.resp.Data.Children {
		_, err := writer.Write([]byte(child.Data.Title + "\n" + child.Data.URL + "\n"))
		if err != nil {
			return err
		}
	}
	return nil
}
