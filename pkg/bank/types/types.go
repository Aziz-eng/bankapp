package types

// Payment is informed about payments.
type Payment struct {
	ID int
	Amount Money
}


//Money it is showed minimum units of money.
type Money int64

// Currency  is showed cod of valute.
type Currency string

// Cod of valyte
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

// PAN showed a number of card
type PAN string

// Card showed information about payment
type Card struct{
	ID  		int
	PAN 		PAN
	Balance  	Money
	Currency  	Currency
	Color  		string
	Name  		string
	Active 		bool
	MinBalance	Money
}
// Payments source it doing somethink
type PaymentSource struct {
	Type string // 'card'
	Number PAN //номер вида '5555 xxxx xxxx 2525'
	Balance Money
}