package main

import (
	"context"
	"os"
	"sort"

	"github.com/LeoAdamek/monzo"
	"github.com/urfave/cli"
)

func topTransactions(ct *cli.Context) error {

	ctx := context.Background()
	token := os.Getenv("MONZO_TOKEN")

	if token == "" {
		log.Fatalln("MONZO_TOKEN environment variable is not set.")
	}

	c := monzo.New(token)

	accounts, err := c.Accounts(ctx)

	if err != nil {
		log.Fatalln("Unable to get accounts: ", err)
	}

	var txns []monzo.Transaction

	for _, account := range accounts {
		transactions, err := c.Transactions(ctx, monzo.ListTransactionsInput{AccountID: account.ID})

		if err != nil {
			log.Printf("Unable to get transactions for account %s: %s", account.ID, err)
			continue
		}

		txns = append(txns, transactions...)
	}

	// Sort transactions
	sort.Sort(monzo.ByValue(txns))

	n := ct.Int("n")

	if n > len(txns) {
		n = len(txns)
	}

	for idx, tx := range txns[0:n] {
		log.Printf("Transaction #%d: %4.2f to %s%s", idx, float64(tx.Amount)/100.0, tx.Merchant.Emoji, tx.Merchant.Name)
	}

	return nil
}
