package main

import (
	"fmt"
)

func main() {
	fmt.Println(makeURLRequest("https://data.ripple.com/v2/accounts/r9Z7PVDTtrNWrW2Uw7BRCs1txUvtWDSLYd/balances?currency=XRP"))
}
