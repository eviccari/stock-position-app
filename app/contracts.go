package app

import "github.com/eviccari/stock-position-app/app/vos"

type DomainService interface {
	FindByItemIDAndFacilityID(itemID string, facilityID int) (sp vos.StockPosVO, err error)
	DoMovement(sm vos.StockMovVO) (sp vos.StockPosVO, err error)
}
