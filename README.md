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
)

func main() {
    token := os.Getenv("MONZO_TOKEN")
    
    if token == "" {
        fmt.Println("Please set the MONZO_TOKEN environment variable")
        os.Exit(1)
    }
    
    m := monzo.New(token)
    
    accounts, err := m.Accounts()
    
    if err != nil {
        fmt.Println("Unable to list accounts:", err)
        os.Exit(1)
    }
    
    transactions, err := m.Transactions(monzo.ListTransactionsInput{AccountID: accounts[0].ID})
    
    if err != nil {
        fmt.Println("Unable to list transactions:", err)
        os.Exit(1)
    }
    
    for _, t := range transactions {
        fmt.Printf("Paid £%4.2d to %s\n", float64(t.Amount)/100.0 , t.Merchant.Name)
    }
}
````


[monzo]: https://monzo.com/

[godoc]: https://godoc.org/LeoAdamek/monzo