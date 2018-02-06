package main

import (
	"github.com/LeoAdamek/monzo"
	"github.com/urfave/cli"
	"os"
	"sort"
)

func topTransactions(ctx *cli.Context) error {

	token := os.Getenv("MONZO_TOKEN")

	if token == "" {
		log.Fatalln("MONZO_TOKEN environment variable is not set.")
	}

	c := monzo.New(token)

	accounts, err := c.Accounts()

	if err != nil {
		log.Fatalln("Unable to get accounts: ", err)
	}

	var txns []monzo.Transaction

	for _, account := range accounts {
		transactions, err := c.Transactions(monzo.ListTransactionsInput{AccountID: account.ID})

		if err != nil {
			log.Printf("Unable to get transactions for account %s: %s", account.ID, err)
			continue
		}

		txns = append(txns, transactions...)
	}

	// Sort transactions
	sort.Sort(monzo.ByValue(txns))

	n := ctx.Int("n")

	if n > len(txns) {
		n = len(txns)
	}

	for idx, tx := range txns[0:n] {
		log.Printf("Transaction #%d: %4.2f to %s", idx, float64(tx.Amount)/100.0, tx.Merchant.Name)
	}

	return nil
}
