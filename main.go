package main

func main() {
	//publicKey1, _ := CreateKeys()
	//publicKey2, _ := CreateKeys()
	//publicKey1 := "GB5FUITDOLTZDQ2YDZBRRDIHE56RLD5TXRSI4FESFN2E74EBLLOHE6YN"
	seed := "SCJW4ZBMGKTORVSOLWXBWCCLO2WHW4VI5UCCQZMXAVXYP7M4IAJ2HYIH"
	publicKey2 := "GDJL7KR3RQ3XH3BSZS3RVSQCPEAA7O4PJKHS55VV7AZF27OJL7QQNPIQ"

	//CreateTestAccount(publicKey1)
	//pub: GB5FUITDOLTZDQ2YDZBRRDIHE56RLD5TXRSI4FESFN2E74EBLLOHE6YN
	//seed: SCJW4ZBMGKTORVSOLWXBWCCLO2WHW4VI5UCCQZMXAVXYP7M4IAJ2HYIH

	//CreateTestAccount(publicKey2)
	//pub: GDJL7KR3RQ3XH3BSZS3RVSQCPEAA7O4PJKHS55VV7AZF27OJL7QQNPIQ
	//seed: SCSJFTW75JRUKTGPABTGNHKKLR2QTANJPID4RHDVZWRFTSFPGHY2GIXJ

	// GetAccountInfo(publicKey1)
	// GetAccountInfo(publicKey2)

	SendPayment(seed, publicKey2)
	ReceivePayment(publicKey2)

	GetAccountInfo(publicKey1)
	GetAccountInfo(publicKey2)
}
