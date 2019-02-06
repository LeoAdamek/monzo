package monzo

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

// FeedEntryType denotes the type of the feed entry.
type FeedEntryType string

const (
	// BasicFeedEntry is the most basic (and currently, only) feed entry type.
	BasicFeedEntry FeedEntryType = "basic"
)

// FeedEntry represents an entry into the user's account feed
type FeedEntry struct {
	AccountID       string        `url:"account_id"`
	Type            FeedEntryType `url:"type"`
	URL             string        `url:"url"`
	Title           string        `url:"params[title]"`
	Body            string        `url:"params[body]"`
	ImageURL        string        `url:"params[image_url]"`
	BackgroundColor string        `url:"params[background_color]"`
	BodyColor       string        `url:"params[body_color]"`
	TitleColor      string        `url:"params[title_color]"`
}

// AddFeedItem adds a new item to the user's account feed
func (c Client) AddFeedItem(ctx context.Context, item FeedEntry) error {

	reqURL := *baseURL
	reqURL.Path = "/feed"

	values, err := query.Values(item)

	if err != nil {
		return err
	}

	body, err := url.QueryUnescape(values.Encode())

	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, reqURL.String(), bytes.NewBufferString(body))

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp := make(map[string]interface{})

	fmt.Println(body)

	err = c.json(ctx, req, &resp)

	return err
}
