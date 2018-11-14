package main

import (
	//"bufio"

	"encoding/json"
	//"os"
)

type jBalance struct {
	Currency string `json:"currency"`
	Value    string `json:"value"`
}

type jResponse struct {
	Result      string     `json:"result"`
	LedgerIndex int        `json:"ledger_index"`
	Limit       int        `json:"limit"`
	Balances    []jBalance `json:"balances"`
}

func parseXRPBalance(rawString string) string {

	//loop code 5 times
	//for i := 1; i < 6; i++ {

	//make https request to retrieve JSON string
	//s := MakeURLRequest("https://data.ripple.com/v2/accounts/r9Z7PVDTtrNWrW2Uw7BRCs1txUvtWDSLYd/balances?")

	//parse JSON string to go variables
	jsonBytes := []byte(rawString)
	jRespObj := jResponse{}
	err := json.Unmarshal(jsonBytes, &jRespObj)
	if err != nil {
		panic(err)
	}

	//report query parameters
	jBalObj := jBalance{}
	for _, element := range jRespObj.Balances {
		jBalObj = element
		//fmt.Printf("Wallet has %s %s\n", jBalObj.Value, jBalObj.Currency)
	}

	//wait for user response
	//fmt.Print("Press 'Enter' to continue...")
	//bufio.NewReader(os.Stdin).ReadBytes('\n')
	//}
	return jBalObj.Value
}
