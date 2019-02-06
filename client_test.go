package monzo

import (
	"context"
	"os"
	"testing"
	"time"
)

var client *Client

var ctx = context.TODO()

func init() {
	token := os.Getenv("MONZO_TOKEN")

	if token == "" {
		panic("MONZO_TOKEN is not set")
	}

	client = New(token)
}

func TestClient_Accounts(t *testing.T) {
	accounts, err := client.Accounts(ctx)

	if err != nil {
		t.Error("Got error listing accounts:", err)
		t.FailNow()
		return
	}

	if len(accounts) == 0 {
		t.Fail()
	}
}

func TestClient_Balance(t *testing.T) {
	accounts, err := client.Accounts(ctx)

	if err != nil {
		t.FailNow()
	}

	b, err := client.Balance(ctx, accounts[0].ID)

	if err != nil {
		t.FailNow()
	}

	t.Log(b)
}

func TestClient_Transactions(t *testing.T) {
	accounts, err := client.Accounts(ctx)

	if err != nil {
		t.FailNow()
	}

	after := time.Now().Truncate(160 * time.Hour)
	before := after.Add(24 * time.Hour)

	transactions, err := client.Transactions(ctx, ListTransactionsInput{
		AccountID: accounts[0].ID,
		Pagination: Pagination{
			Since:  &after,
			Before: &before,
			Limit:  5,
		},
	})

	if err != nil {
		t.Error("Unable to get transactions:", err)
		t.Fail()
		return
	}

	if len(transactions) == 0 {
		t.Fail()
	}
}

func TestClient_GetTransaction(t *testing.T) {
	accounts, err := client.Accounts(ctx)

	if err != nil {
		t.FailNow()
	}

	transactions, err := client.Transactions(ctx, ListTransactionsInput{AccountID: accounts[0].ID})

	if err != nil {
		t.FailNow()
	}

	tid := transactions[0].ID

	txn, err := client.GetTransaction(ctx, tid)

	if err != nil {
		t.FailNow()
	}

	t.Log(txn)
}

func TestClient_AddFeedItem(t *testing.T) {

	// Skip this for short as it will push an item to the user's feed.
	if testing.Short() {
		t.SkipNow()
	}

	accounts, err := client.Accounts(ctx)

	if err != nil {
		t.FailNow()
	}

	aid := accounts[0].ID

	item := FeedEntry{
		AccountID:       aid,
		Type:            BasicFeedEntry,
		URL:             "about:blank",
		ImageURL:        "https://www.placecage.com/32/32",
		Title:           "Monzo API Test",
		Body:            "Test API entry",
		BackgroundColor: "#000000",
		BodyColor:       "#AAAAAA",
		TitleColor:      "#FFFFFF",
	}

	err = client.AddFeedItem(ctx, item)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
