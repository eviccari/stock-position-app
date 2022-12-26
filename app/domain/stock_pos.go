package domain

import (
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
)

type StockPos struct {
	id            string
	itemID        string
	facilityID    uint
	onHandQty     int
	unavlQty      int
	transactionID string
	positionDate  time.Time
}

func StartStockPos(itemID string, facilityID uint, transactionID string) StockPos {
	return StockPos{
		id:            uuid.NewString(),
		itemID:        itemID,
		facilityID:    facilityID,
		onHandQty:     0,
		unavlQty:      0,
		transactionID: transactionID,
		positionDate:  time.Now(),
	}
}

func (sp *StockPos) DoMovement(sm StockMov) (errorList []error) {
	before := sp.Clone()
	sp.onHandQty += sm.toAvailQty - sm.fromAvailQty
	sp.unavlQty += sm.toUnavlQty - sm.fromUnavlQty
	sp.positionDate = time.Now()
	sp.transactionID = sm.transactionID
	logTransaction(*before, *sp, sm)

	return ApplyPolicies(*sp, GetPoliciesID()...)
}

func (sp *StockPos) Clone() *StockPos {
	return &StockPos{
		id:            sp.id,
		itemID:        sp.itemID,
		facilityID:    sp.facilityID,
		onHandQty:     sp.onHandQty,
		unavlQty:      sp.unavlQty,
		transactionID: sp.transactionID,
		positionDate:  sp.positionDate,
	}
}

func logTransaction(before, after StockPos, sm StockMov) {
	smL := struct {
		TransactionID string `json:"transaction_id"`
		ItemID        string `json:"item"`
		Facility      uint   `json:"facility"`
		FromAvail     int    `json:"from_avail"`
		ToAvail       int    `json:"to_avail"`
		FromUnavl     int    `json:"from_unavl"`
		ToUnavl       int    `json:"to_unavl"`
	}{
		sm.transactionID,
		sm.itemID,
		sm.facilityID,
		sm.fromAvailQty,
		sm.toAvailQty,
		sm.fromUnavlQty,
		sm.toUnavlQty,
	}

	spL := struct {
		TransactionID string `json:"transaction_id"`
		ItemID        string `json:"item"`
		FacilityID    uint   `json:"facility"`
		BeforeAvail   int    `json:"before_avail"`
		AfterAvail    int    `json:"after_avail"`
		BeforeUnavl   int    `json:"before_unavl"`
		AfterUnavl    int    `json:"after_unavl"`
	}{
		before.transactionID,
		before.itemID,
		before.facilityID,
		before.onHandQty,
		after.onHandQty,
		before.unavlQty,
		after.unavlQty,
	}

	j, _ := json.Marshal(smL)
	log.Printf("STK_MOV::%s\n", string(j))

	j, _ = json.Marshal(spL)
	log.Printf("STK_POS::%s\n", string(j))
}
