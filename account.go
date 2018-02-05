package monzo

import (
    "time"
    "net/http"
)

// Account represents an account
type Account struct {
    ID string `json:"id"`
    Description string `json:"description"`
    CreatedAt time.Time `json:"created"`
}

// Accounts gets a list of the user's accounts
func (c Client) Accounts() ([]Account, error) {
    
    reqURL := *baseURL
    reqURL.Path = "/accounts"
    
    req := &http.Request{
        Method: http.MethodGet,
        URL: &reqURL,
    }
    
    var response struct {
        Accounts []Account
    }
    
    err := c.json(req, &response)
    
    return response.Accounts, err
}