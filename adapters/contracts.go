package adapters

import "github.com/eviccari/stock-position-app/app/vos"

type Repository interface {
	Create(vo vos.StockPosVO) (err error)
	Update(vo vos.StockPosVO, constraintKey string) (updatedRowCount int64, err error)
	FindByItemIDAndFacilityID(itemID string, facilityID uint) (vo vos.StockPosVO, err error)
}
