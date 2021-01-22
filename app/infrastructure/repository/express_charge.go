package repository

import (
	iRep "github.com/mYahagi/go-api-sample/app/domain/repository"
	"github.com/mYahagi/go-api-sample/app/infrastructure/model"
)

type ExpressChargeRepository struct{}

func (r ExpressChargeRepository) Find() (*model.ExpressCharge, error) {
	return &model.ExpressCharge{
		Charge: 100,
	}, nil
}

func NewExpressChargeRepository() iRep.IExpressChargeRepository {
	return ExpressChargeRepository{}
}
