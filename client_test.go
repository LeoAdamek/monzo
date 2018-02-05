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
    
    t.Log(accounts)
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
    
    t.Log(transactions)
}