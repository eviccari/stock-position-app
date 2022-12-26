package domain

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/eviccari/stock-position-app/utils"
)

var policies = map[string]func(sp StockPos) (err error){}

func StockPositionKeysAreRequired(sp StockPos) (err error) {
	if utils.IsEmptyString(sp.itemID) || sp.facilityID == 0 {
		m := "item_id and facility_id are required"
		logErrorMessage(sp.transactionID, m)
		err = errors.New(m)
	}
	return
}

func AvailQtyCannotBeLessThanZero(sp StockPos) (err error) {
	if sp.onHandQty < 0 {
		m := "avail_qty cannot be less than zero"
		logErrorMessage(sp.transactionID, m)
		err = errors.New(m)
	}
	return
}

func UnavlQtyCannotBeLessThanZero(sp StockPos) (err error) {
	if sp.unavlQty < 0 {
		m := "unavl_qty cannot be less than zero"
		logErrorMessage(sp.transactionID, m)
		err = errors.New(m)
	}
	return
}

func TransactionIDIsRequired(sp StockPos) (err error) {
	if utils.IsEmptyString(sp.transactionID) {
		m := "transaction_id is required"
		logErrorMessage(sp.transactionID, m)
		err = errors.New(m)
	}
	return
}

func ApplyPolicies(sp StockPos, policyNames ...string) (errorList []error) {
	for _, pn := range policyNames {
		f := policies[pn]
		if err := f(sp); err != nil {
			errorList = append(errorList, err)
		}
	}
	return
}

func init() {
	policies["StockPositionKeysAreRequired"] = StockPositionKeysAreRequired
	policies["AvailQtyCannotBeLessThanZero"] = AvailQtyCannotBeLessThanZero
	policies["UnavlQtyCannotBeLessThanZero"] = UnavlQtyCannotBeLessThanZero
	policies["TransactionIDIsRequired"] = TransactionIDIsRequired
}

func GetPoliciesID() (policiesID []string) {
	for p := range policies {
		policiesID = append(policiesID, p)
	}
	return
}

func logErrorMessage(transactionID, message string) {
	m := struct {
		TransactionID string `json:"transaction_id"`
		Message       string `json:"message"`
	}{transactionID, message}

	j, _ := json.Marshal(m)
	log.Printf("STK_ERR::%s\n", string(j))
}
