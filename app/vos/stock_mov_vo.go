package vos

type StockMovVO struct {
	ID                     string
	ItemID                 string
	FacilityID             uint
	FromAvailQty           int
	ToAvailQty             int
	FromUnavlQty           int
	ToUnavlQty             int
	TransactionDescription string
	TransactionID          string
}
