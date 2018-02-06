package monzo

import (
	"os"
	"testing"
)

var client *Client

func init() {
	token := os.Getenv("MONZO_TOKEN")

	if token == "" {
		panic("MONZO_TOKEN is not set")
	}

	client = New(token)
}

func TestClient_Accounts(t *testing.T) {
	accounts, err := client.Accounts()

	if err != nil {
		t.Error("Got error listing accounts:", err)
		t.FailNow()
		return
	}

	if len(accounts) == 0 {
		t.Fail()
	}
}

func TestClient_Transactions(t *testing.T) {
	accounts, err := client.Accounts()

	if err != nil {
		t.FailNow()
	}

	transactions, err := client.Transactions(ListTransactionsInput{
		AccountID: accounts[0].ID,
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

func TestClient_AddFeedItem(t *testing.T) {
	accounts, err := client.Accounts()

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

	err = client.AddFeedItem(item)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
