package card

import (
	"bank/pkg/bank/types"
	// "fmt"
)

//Withdraw is function of monitoring status card for using
// func Withdraw(card types.Card, amount types.Money) types.Card {

// 	if card.Active == true {
// 		if amount > 0 {
// 			if amount < card.Balance {
// 				if amount < 20_000_00 {
// 					card.Balance = card.Balance - amount
// 					if amount > 20_000_00 {
// 						fmt.Println(card.Balance)
// 					}
// 				}
// 				if amount > card.Balance {
// 					fmt.Println(card.Balance)
// 				}
// 			}
// 			if amount < 0 {
// 				fmt.Println(card.Balance)
// 			}
// 		}
// 	}
// 	if card.Active == false {
// 		fmt.Println()
// 	}

// 	return card
// }

// const amountLimit = 50_000_00

// //Deposit function for including money
// func Deposit(card *types.Card, amount types.Money) {
// 	if amount < 0 {
// 		return
// 	}
// 	if !card.Active {
// 		return
// 	}
// 	if amount > amountLimit {
// 		return
// 	}
// 	(*card).Balance = card.Balance + amount

// }

// const bonusLimit = 500
// AddBonus is function for enrollment to balans
// func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {

// 	if card.Active == false {
// 		return
// 	}
// 	if card.MinBalance <= 0 {
// 		return
// 	}
// 	if card.Balance <= 0 {
// 		return
// 	}

// 	bonus := ((int(card.MinBalance) * percent / 100 * daysInMonth) / daysInYear)
// 	if bonus > 500000 || bonus < 0 {
// 		return
// 	}
// 	card.MinBalance = card.Balance
// 	(*card).Balance += types.Money(bonus)

// }

// IssueCard is Issue a new card
// func IssueCard(currency types.Currency, color string, name string) types.Card {
// 	exitCard := types.Card{
// 		ID:       1000,
// 		PAN:      "5058 xxxx xxxx 0001",
// 		Balance:  0,
// 		Currency: types.TJS,
// 		Color:    color,
// 		Name:     name,
// 		Active:   true,
// 	}
// 	return exitCard
// }
// Total for sum of balanses
func Total(cards []types.Card) types.Money{
	
		var sum types.Money
	for _, card := range cards {
		sum += card.Balance 
		
	}
	// fmt.Println(Total(cards))
return sum
}
// Payment sources.
func PaymentSources(cards []types.Card) []types.PaymentSource {
		paymentSources := []types.PaymentSource {}
		for _, card := range cards {
			if card.Balance <= 0 {
				continue
			}
			if !card.Active  {
			
			continue
			}
			
			paymentSources = append(paymentSources, types.PaymentSource{ Number: card.PAN, Balance: card.Balance,})
		}
	return paymentSources
}