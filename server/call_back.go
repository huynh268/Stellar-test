package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type PaymentInfo struct {
	Amount        string `json:"amount"`
	AssetCode     string `json:"asset_code"`
	AssetIssuer   string `json:"asset_issuer"`
	Data          string `json:"data"`
	From          string `json:"from"`
	ID            string `json:"id"`
	Memo          string `json:"memo"`
	MemoType      string `json:"memo_type"`
	Route         string `json:"route"`
	TransactionID string `json:"transaction_id"`
}

var (
	flagPort = flag.String("port", "8005", "Bridge server callbacks running on port 8005!")
)

//ReceivePayment get the callbacks from bridge server
func ReceivePayment(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
		}
		fmt.Println(string(body))

		var p PaymentInfo
		json.Unmarshal(body, &p)
		//fmt.Println("Amount:" + p.From)
		fmt.Fprint(w, "POST done!")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/receive", ReceivePayment)
	log.Printf("listening on port %s", *flagPort)
	log.Fatal(http.ListenAndServe(":"+*flagPort, mux))

}
