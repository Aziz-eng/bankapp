package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

// func ExampleWithdraw_positive () {
// 	result:= Withdraw(types.Card{Balance: 20_000_00, Active: true}, 10_000_00)
// fmt.Println(result.Balance)
// // Output: 1000000
// }

// func ExampleWithdraw_noMoney () {
// 	result:= Withdraw(types.Card{Balance: 0, Active: true}, 10_000_00)
// fmt.Println(result.Balance)
// // Output: 0
// }

// func ExampleWithdraw_inactive ()  {
// 	result:= Withdraw(types.Card{Balance: 20_000_00, Active: false}, 10_000_00)
// 	fmt.Println(result.Balance)
// 	// Output: 2000000
// }

// func ExampleWithdraw_limit ()  {
// 	result:= Withdraw(types.Card{Balance: 20_000_00, Active: true}, 21_000_00)
// 	fmt.Println(result.Balance)
// 	// Output: 2000000

// }
// // Testing Deposit
// func ExampleDeposit_nomarl ()  {
// 	result:= types.Card {Balance: 20_000_00, Active: true}
// 	Deposit(&result, 10_000_00)
// 	fmt.Println(result.Balance)
//  	// Output: 3000000
// }
// func ExampleDeposit_inactive()  {
// 	result:= types.Card {Balance: 20_000_00, Active: false}
// 	Deposit(&result, 10_000_00)
// 	fmt.Println(result.Balance)
// 	// Output: 2000000
// }
// func ExampleDeposit_limit ()  {
// result:= types.Card {Balance: 20_000_00, Active: true}
// 	Deposit(&result, 51_000_00)
// 	fmt.Println(result.Balance)
// 	// Output: 2000000
// }

// func ExampleAddBonus_positive() {
// 	result := types.Card{Balance: 2460, Active: true}
// 	AddBonus(&result, 3, 30, 365)
// 	fmt.Println(result.Balance)

// 	// Output: 2466
// }

// func ExampleAddBonus_positive() {
// 	card := types.Card{MinBalance: 10_000_00, Balance: 12_000_00, Active: true}
// 	AddBonus(&card, 3, 30, 365)
// 	fmt.Println(card.Balance)
// 	// Output: 1202465
// }
// func ExampleAddBonus_inactive() {
// 	result := types.Card{MinBalance: 10000, Active: false}
// 	AddBonus(&result, 3, 30, 365)
// 	fmt.Println(result.MinBalance)

// 	//Output: 10000
// }

// func ExampleAddBonus_nomoney() {
// 	result := types.Card{MinBalance: 0, Balance: -1, Active: true}
// 	AddBonus(&result, 3, 30, 365)
// 	fmt.Println(result.MinBalance)

// 	//Output: 0
// }
// func ExampleAddBonus_limit() {
// 	result := types.Card{MinBalance: 10_000_000_00, Active: true}
// 	AddBonus(&result, 3, 30, 365)
// 	fmt.Println(result.MinBalance)

// 	//Output: 1000000000
// }

func ExampleTotal() {
	cards := []types.Card{
		types.Card{
			Balance: 10_000_00,
			Active:  true,
		},
		types.Card{
			Balance: 10_000_00,
			Active:  true,
		},
	}
	fmt.Println(Total(cards))
	// Output: 2000000
}