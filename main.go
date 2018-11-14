package main

import (
	"fmt"
)

func main() {

	//get XRP balance
	pubKey := txtFileToString("./PublicKeys.txt")
	rawString := makeURLRequest("https://data.ripple.com/v2/accounts/" + pubKey + "/balances?currency=XRP")
	balance := parseXRPBalance(rawString)
	fmt.Println(balance)

}
