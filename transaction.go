package monzo

import (
	"net/http"
	"strconv"
	"time"
)

// Transaction represents a transaction
type Transaction struct {
	ID string `json:"id"`

	CreatedAt time.Time `json:"created"`

	Amount int `json:"amount"`

	// Account balance after this transaction. Used only for Mondo prepaid cards.
	AccountBalance int `json:"account_balance"`

	Currency string `json:"currency"`

	Description string `json:"description"`

	// Notes against the transaction
	Notes string `json:"notes"`

	// SettledAt may be an empty string.
	// TODO: Handle this field.
	// SettledAt time.Time `json:"settled"`

	// Custom transaction metadata.
	Metadata TransactionMetadata

	// The merchant or other counter-party involved in the transaction.
	Merchant Merchant
}

// TransactionMetadata is a set of custom key/value pairs assigned to a transaction.
type TransactionMetadata map[string]string

// Merchant represents a Merchant
type Merchant struct {
	CreatedAt time.Time `json:"created"`
	GroupID   string    `json:"group_id"`
	ID        string    `json:"id"`

	Address struct {
		Address   string  `json:"address"`
		City      string  `json:"city"`
		Country   string  `json:"country"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Postcode  string  `json:"postcode"`
		Region    string  `json:"region"`
	}

	Logo     string `json:"logo"`
	Emoji    string `json:"emoji"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

// ListTransactionsInput is used for the parameters of a transaction listing request
type ListTransactionsInput struct {
	AccountID string `json:"account_id"`
	Pagination
}

// GetTransaction gets a single transaction
func (c Client) GetTransaction(transactionID string) (Transaction, error) {

	reqURL := *baseURL
	reqURL.Path = "/transactions/" + transactionID

	q := reqURL.Query()
	q.Set("expand[]", "merchant")
	reqURL.RawQuery = q.Encode()

	req := &http.Request{
		Method: http.MethodGet,
		URL:    &reqURL,
	}

	var transaction Transaction

	err := c.json(req, &transaction)

	return transaction, err
}

// Transactions lists a set of transactions
//
// input will define how those transactions are used
func (c Client) Transactions(input ListTransactionsInput) ([]Transaction, error) {

	reqURL := *baseURL
	reqURL.Path = "/transactions"
	q := reqURL.Query()

	q.Set("account_id", input.AccountID)
	q.Set("limit", strconv.Itoa(100))
	q.Add("expand[]", "merchant")

	if input.Since != nil {
		q.Set("since", input.Since.Format(time.RFC3339))
	}

	if input.Before != nil {
		q.Set("before", input.Before.Format(time.RFC3339))
	}

	if input.Limit > 0 && input.Limit <= 100 {
		q.Set("limit", strconv.Itoa(input.Limit))
	}

	reqURL.RawQuery = q.Encode()

	req := &http.Request{
		Method: http.MethodGet,
		URL:    &reqURL,
	}

	var response struct {
		Transactions []Transaction
	}

	err := c.json(req, &response)

	return response.Transactions, err
}

// ByValue allows for sorting of transactions by their value.
//
// Example:
//   var t []Transaction
//   sort.Sort(ByValue(t))
type ByValue []Transaction

//
func (b ByValue) Len() int {
	return len(b)
}

//
func (b ByValue) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

//
func (b ByValue) Less(i, j int) bool {
	return b[i].Amount < b[j].Amount
}
