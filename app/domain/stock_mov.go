package domain

type StockMov struct {
	id                     string
	itemID                 string
	facilityID             uint
	fromAvailQty           int
	toAvailQty             int
	fromUnavlQty           int
	toUnavlQty             int
	transactionDescription string
	transactionID          string
}
