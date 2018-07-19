package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/stellar/go/clients/horizon"
	"github.com/stellar/go/keypair"
)

//CreateKeys creates public key and seed
func CreateKeys() (string, string) {
	pair, err := keypair.Random()
	if err != nil {
		log.Fatal(err)
	}

	publicKey := pair.Address()
	//GCSGYI3B4WGZTNQV5FKNMFNE5MWQDRN6JXIXPEIL3VF54RJ3XLZIPPRH

	seed := pair.Seed()
	//SBW3DHJCLTPTTFPFCBZBVF335ZT3SS6KU7TM7H6GXO7ACUZ3ZYSXTAOG

	log.Println("Public key: ", publicKey)
	log.Println("Seed: ", seed)

	return publicKey, seed
}

//CreateTestAccount creates a test account
func CreateTestAccount(publicKey string) {
	resp, err := http.Get("https://friendbot.stellar.org/?addr=" + publicKey)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}

//GetAccountInfo gets info about the account just created
func GetAccountInfo(publicKey string) {
	account, err := horizon.DefaultTestNetClient.LoadAccount(publicKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Balances for account:", publicKey)

	for _, balance := range account.Balances {
		log.Println(balance)
	}
}
