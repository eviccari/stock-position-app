package services

import (
	"errors"

	"github.com/eviccari/stock-position-app/adapters"
	"github.com/eviccari/stock-position-app/app/domain"
	"github.com/eviccari/stock-position-app/app/vos"
	"github.com/eviccari/stock-position-app/utils"
)

type StockService struct {
	repo adapters.Repository
}

func NewStockService(repo adapters.Repository) *StockService {
	return &StockService{
		repo: repo,
	}
}

func (ss *StockService) Update(sm vos.StockMovVO) (spVO vos.StockPosVO, err error) {
	spVO, err = ss.repo.FindByItemIDAndFacilityID(sm.ItemID, sm.FacilityID)
	if err != nil {
		return
	}

	var spDomain = domain.StockPos{}
	if utils.IsEmptyString(spVO.ID) {
		spDomain = domain.StartStockPos(sm.ItemID, sm.FacilityID, sm.TransactionID)
		spVO = domain.BuildStockPosVO(spDomain)
		ss.repo.Create(spVO)
	} else {
		spDomain = domain.BuildStockPos(spVO)
	}

	constraintKey := spVO.TransactionID

	errorList := spDomain.DoMovement(domain.BuildStockMov(sm))
	if len(errorList) > 0 {
		err = errors.New(handleErrorList(errorList))
		return
	}

	spVO = domain.BuildStockPosVO(spDomain)
	ss.repo.Update(spVO, constraintKey)

	return
}

func (ss *StockService) FindByItemIDAndFacilityID(itemID string, facilityID int) (sp vos.StockPosVO, err error) {
	return ss.repo.FindByItemIDAndFacilityID(itemID, sp.FacilityID)
}

func handleErrorList(el []error) (s string) {
	for _, e := range el {
		s += e.Error() + ", "
	}
	return s
}
