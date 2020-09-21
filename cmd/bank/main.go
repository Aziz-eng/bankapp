package main

import (
	"bank/pkg/bank/card"
	"bank/pkg/bank/types"
	"fmt"
)

func main() {
	// cards:= []types.Card {
	// 	Balance: 100000,
	// 	Active: true,
	// }
	// fmt.Println (cards)
	cards := []types.Card{
		// {

		// 	PAN:     string,
		// 	Balance: Money,
		// 	Active:  true,
		// },
		{

			PAN:     "2525 xxxx xxxx 0002",
			Balance: 25000,
			Active:  true,
		},
	}
	paymentSources := card.PaymentSources(cards)
	for _, v := range paymentSources {
		fmt.Println(v.Number)
	}

}
