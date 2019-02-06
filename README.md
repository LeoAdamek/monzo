[Monzo](.)
=====

Golang API client for [Monzo][monzo]

Features
--------

* (ONE) external dependency.
* Overridable `net/http.Client`
* Overridable logging
* Decent test coverage (ish)
* Built-in helpers

Documentation
------------- 

Check [godoc][godoc] for full documentation.

## A simple example.

````go
package main
import (
    "github.com/LeoAdamek/monzo"
    "os"
    "fmt"
    "context"
)

func main() {
    token := os.Getenv("MONZO_TOKEN")
    ctx := context.WithTimeout(time.Second, context.Background())
    
    if token == "" {
        fmt.Println("Please set the MONZO_TOKEN environment variable")
        os.Exit(1)
    }
    
    m := monzo.New(token)
    
    accounts, err := m.Accounts(ctx)
    
    if err != nil {
        fmt.Println("Unable to list accounts:", err)
        os.Exit(1)
    }
    
    transactions, err := m.Transactions(ctx, monzo.ListTransactionsInput{AccountID: accounts[0].ID})
    
    if err != nil {
        fmt.Println("Unable to list transactions:", err)
        os.Exit(1)
    }
    
    for _, t := range transactions {
        fmt.Printf("Paid £%s to %s\n", t.Amount.String() , t.Merchant.Name)
    }
}
````


[monzo]: https://monzo.com/

[godoc]: https://godoc.org/LeoAdamek/monzo