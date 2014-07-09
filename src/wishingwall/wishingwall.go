package main

import (
	"fmt"
	//"github.com/Wishing-Wall/btcrpcclient"
	//"log"
	//. "protocol"
)

func main() {
	fmt.Println("Hello world!")
	/***
		client, err := btcrpcclient.New(&btcrpcclient.ConnConfig{
			HttpPostMode: true,
			DisableTLS:   true,
			Host:         "127.0.0.1:18001",
			User:         "dogecoin",
			Pass:         "DL2rnNd5KnA5HPFW9J84upxyXXRN9yoe3D",
		}, nil)
		if err != nil {
			log.Fatalf("error creating new btc client: %v", err)
		}

		// list accounts
		accounts, err := client.ListAccounts()
		if err != nil {
			log.Fatalf("error listing accounts: %v", err)
		}
		// iterate over accounts (map[string]btcutil.Amount) and write to stdout
		for label, amount := range accounts {
			log.Printf("%s: %s", label, amount)
		}
	***/
	//Block()
}
