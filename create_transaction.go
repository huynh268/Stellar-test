package main

import (
	"context"
	"fmt"

	"github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
)

//SendPayment makes a transaction between source to destination
func SendPayment(source, destination string) {
	//source := "SCZANGBA5YHTNYVVV4C3U252E2B6P6F5T3U6MM63WBSBZATAQI3EBTQ4"
	//destination := "GA2C5RFPE6GCKMY3US5PAB6UZLKIGSPIUKSLRB6Q723BM2OARMDUYEJ5"

	// Make sure destination account exists
	if _, err := horizon.DefaultTestNetClient.LoadAccount(destination); err != nil {
		panic(err)
	}

	//passphrase := network.TestNetworkPassphrase

	tx, err := build.Transaction(
		build.TestNetwork,
		build.SourceAccount{source},
		build.AutoSequence{horizon.DefaultTestNetClient},
		build.MemoText{"Test Transaction"},
		build.Payment(
			build.Destination{destination},
			build.NativeAmount{"10"},
		),
	)

	if err != nil {
		panic(err)
	}

	// Sign the transaction to prove you are actually the person sending it.
	txe, err := tx.Sign(source)
	if err != nil {
		panic(err)
	}

	txeB64, err := txe.Base64()
	if err != nil {
		panic(err)
	}

	// And finally, send it off to Stellar!
	resp, err := horizon.DefaultTestNetClient.SubmitTransaction(txeB64)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successful Transaction:")
	fmt.Println("Ledger:", resp.Ledger)
	fmt.Println("Hash:", resp.Hash)
}

//ReceivePayment checks the payment
func ReceivePayment(address string) {
	ctx := context.Background()

	cursor := horizon.Cursor("now")

	fmt.Println("Waiting for a payment...")

	err := horizon.DefaultTestNetClient.StreamPayments(ctx, address, &cursor, func(payment horizon.Payment) {
		fmt.Println("Payment type:", payment.Type)
		fmt.Println("Payment Paging Token:", payment.PagingToken)
		fmt.Println("Payment From:", payment.From)
		fmt.Println("Payment To:", payment.To)
		fmt.Println("Payment Asset Type:", payment.AssetType)
		fmt.Println("Payment Asset Code:", payment.AssetCode)
		fmt.Println("Payment Asset Issuer:", payment.AssetIssuer)
		fmt.Println("Payment Amount:", payment.Amount)
		fmt.Println("Payment Memo Type:", payment.Memo.Type)
		fmt.Println("Payment Memo:", payment.Memo.Value)
	})

	if err != nil {
		panic(err)
	}
}
