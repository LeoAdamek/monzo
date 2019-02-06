// This is a sample command line application to demonstrate the Mozno API
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	logging "log"
	"net/http"
	"os"
	"time"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/urfave/cli"
)

var app *cli.App
var log *logging.Logger

type tokenResponse struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"expires_at"`
}

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

	app.Before = setupToken

	log = logging.New(app.Writer, "cli ", logging.LstdFlags)
}

func setupToken(c *cli.Context) error {

	if token := os.Getenv("MONZO_TOKEN"); token == "" {

		saveToken := false

		tokenPath, err := homedir.Expand("~/.monzo_cli")

		if err != nil {
			log.Fatalln("Unable to determine home directory")
		}

		tokenFile, err := os.OpenFile(tokenPath, os.O_RDWR|os.O_CREATE, 0600)

		if err == nil {
			saveToken = true
			defer tokenFile.Close()

			dec := json.NewDecoder(tokenFile)

			token := tokenResponse{}

			if err = dec.Decode(&token); err == nil {
				if token.Expires.After(time.Now()) {
					os.Setenv("MONZO_TOKEN", token.Token)
					return nil
				}
			}

		} else {
			log.Println("Unable to store monzo token. You will need to re-authenticate each run!")
			log.Println("Error detail: ", err)
		}

		resp, err := http.Get("https://monzo-auth.adamek.io/new")

		if err != nil {
			return err
		}

		content, err := ioutil.ReadAll(resp.Body)

		data := make(map[string]string)
		err = json.Unmarshal(content, &data)

		if err != nil {
			return err
		}

		fmt.Println("Please visit this URL to authorize:", data["login_url"])

		t := time.NewTicker(time.Second)

		for i := 0; i < 300; i++ {
			<-t.C

			resp, err := http.Get("https://monzo-auth.adamek.io/token?id=" + data["id"])

			if err != nil {
				return err
			}

			if resp.StatusCode == http.StatusOK {
				content, err := ioutil.ReadAll(resp.Body)

				if err != nil {
					return err
				}

				data := tokenResponse{}

				err = json.Unmarshal(content, &data)

				if err != nil {
					return err
				}

				os.Setenv("MONZO_TOKEN", data.Token)

				if saveToken {
					tokenFile.Truncate(0)

					enc := json.NewEncoder(tokenFile)
					enc.Encode(data)

				}

				return nil
			}
		}
		// It's been 5 minutes, give up.
		return errors.New("token request timed out")
	}

	return nil
}

func main() {
	app.Run(os.Args)
}
