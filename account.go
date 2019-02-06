package monzo

import (
	"context"
	"net/http"
	"time"
)

// Account represents an account
type Account struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created"`
}

// Balance represents an account balance
type Balance struct {
	Balance    int64  `json:"balance"`
	SpendToday int64  `json:"spend_today"`
	Currency   string `json:"currency"`
}

// Accounts gets a list of the user's accounts
func (c Client) Accounts(ctx context.Context) ([]Account, error) {

	reqURL := *baseURL
	reqURL.Path = "/accounts"

	req := &http.Request{
		Method: http.MethodGet,
		URL:    &reqURL,
	}

	var response struct {
		Accounts []Account
	}

	err := c.json(ctx, req, &response)

	return response.Accounts, err
}

// Balance gets teh current balance for a given account.
func (c Client) Balance(ctx context.Context, accountID string) (Balance, error) {
	reqURL := *baseURL
	reqURL.Path = "/balance"

	q := reqURL.Query()
	q.Set("account_id", accountID)

	reqURL.RawQuery = q.Encode()

	req := &http.Request{
		Method: http.MethodGet,
		URL:    &reqURL,
	}

	var b Balance

	err := c.json(ctx, req, &b)

	return b, err
}
