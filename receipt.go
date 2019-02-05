package monzo

import "net/http"

// Receipt represents a full itemized receipt.
type Receipt struct {
	ID            string           `json:"id"`
	ExternalID    string           `json:"external_id"`
	TransactionID string           `json:"transaction_id"`
	Total         Money            `json:"total"`
	Currency      string           `json:"currency"`
	Items         []ReceiptItem    `json:"items"`
	Taxes         []ReceiptTax     `json:"taxes"`
	Payments      []ReceiptPayment `json:"payments"`
	Merchant      ReceiptMerchant  `json:"merchant"`
}

// ReceiptItem represents a single item of a receipt
type ReceiptItem struct {
	Description string        `json:"description"`
	Amount      Money         `json:"amount"`
	Currency    string        `json:"currency"`
	Quantity    float64       `json:"quantity"`
	Unit        string        `json:"unit"`
	Tax         Money         `json:"tax"`
	SubItems    []ReceiptItem `json:"sub_items"`
}

// ReceiptTax represents a tax levied on a receipt
type ReceiptTax struct {
	Description string `json:"description"`
	Amount      Money  `json:"amount"`
	Currency    string `json:"currency"`
	TaxNumber   string `json:"tax_number"`
}

// ReceiptPayment represents a payment made against this receipt
type ReceiptPayment struct {
	Type         string `json:"type"`
	Amount       Money  `json:"amount"`
	Currency     string `json:"currency"`
	LastFour     string `json:"last_four"`
	GiftGardType string `json:"gift_card_type"`
}

// ReceiptMerchant represents merchant data within the receipt
type ReceiptMerchant struct {
	Name          string `json:"name"`
	IsOnline      bool   `json:"online"`
	PhoneNumber   string `json:"phone"`
	EmailAddress  string `json:"email"`
	StoreName     string `json:"store_name"`
	StoreAddress  string `json:"store_address"`
	StorePostcode string `json:"store_postcode"`
}

// GetReceipt will retrieve a single Receipt by its `external_id`
//
// @param externalID External ID of a receipt.
func (c Client) GetReceipt(externalID string) (*Receipt, error) {
	reqURL := *baseURL

	reqURL.Path = "/transaction-receipts"
	q := reqURL.Query()

	q.Set("external_id", externalID)

	reqURL.RawQuery = q.Encode()

	req := &http.Request{
		URL:    &reqURL,
		Method: http.MethodGet,
	}

	r := &Receipt{}

	if err := c.json(req, &r); err != nil {
		return nil, err
	}

	return r, nil
}
