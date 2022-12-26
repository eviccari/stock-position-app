package vos

import "time"

type StockPosVO struct {
	ID            string    `json:"id"`
	ItemID        string    `json:"item_id"`
	FacilityID    uint      `json:"facility_id"`
	OnHandQty     int       `json:"on_hand_qty"`
	UnavlQty      int       `json:"unavl_qty"`
	TransactionID string    `json:"transaction_id"`
	PositionDate  time.Time `json:"position_date"`
}
