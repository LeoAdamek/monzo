// This is a sample command line application to demonstrate the Mozno API
package main

import (
	"github.com/urfave/cli"
	logging "log"
	"os"
)

var app *cli.App
var log *logging.Logger

func init() {
	app = cli.NewApp()
	app.Name = "Monzo CLI"
	app.Author = "Leo Adamek <code@breakerofthings.tech>"
	app.Description = "An example application for the Monzo API client"
	app.EnableBashCompletion = true

	app.Commands = []cli.Command{
		{
			Name:   "top",
			Usage:  "Show the top transactions on an account (by largest value)",
			Action: topTransactions,
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "n",
					Value: 10,
					Usage: "Number of transactions to show",
				},
			},
		},
	}

	log = logging.New(app.Writer, "cli ", logging.LstdFlags)
}

func main() {
	app.Run(os.Args)
}
