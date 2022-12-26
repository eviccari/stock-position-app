package domain

import "github.com/eviccari/stock-position-app/app/vos"

func BuildStockPos(vo vos.StockPosVO) StockPos {
	return StockPos{
		id:            vo.ID,
		itemID:        vo.ItemID,
		facilityID:    vo.FacilityID,
		onHandQty:     vo.OnHandQty,
		unavlQty:      vo.UnavlQty,
		transactionID: vo.TransactionID,
		positionDate:  vo.PositionDate,
	}
}

func BuildStockPosVO(sp StockPos) vos.StockPosVO {
	return vos.StockPosVO{
		ID:            sp.id,
		ItemID:        sp.itemID,
		FacilityID:    sp.facilityID,
		OnHandQty:     sp.onHandQty,
		UnavlQty:      sp.unavlQty,
		TransactionID: sp.transactionID,
		PositionDate:  sp.positionDate,
	}
}

func BuildStockMov(vo vos.StockMovVO) StockMov {
	return StockMov{
		id:                     vo.ID,
		itemID:                 vo.ItemID,
		facilityID:             vo.FacilityID,
		fromAvailQty:           vo.FromAvailQty,
		toAvailQty:             vo.ToAvailQty,
		fromUnavlQty:           vo.FromUnavlQty,
		toUnavlQty:             vo.ToUnavlQty,
		transactionDescription: vo.TransactionDescription,
		transactionID:          vo.TransactionID,
	}
}

func BuildStockMovVO(sm StockMov) vos.StockMovVO {
	return vos.StockMovVO{
		ID:                     sm.id,
		ItemID:                 sm.itemID,
		FacilityID:             sm.facilityID,
		FromAvailQty:           sm.fromAvailQty,
		ToAvailQty:             sm.toAvailQty,
		FromUnavlQty:           sm.fromUnavlQty,
		ToUnavlQty:             sm.toUnavlQty,
		TransactionDescription: sm.transactionDescription,
		TransactionID:          sm.transactionID,
	}
}
