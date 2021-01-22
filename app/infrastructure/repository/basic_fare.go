package repository

import (
	iRep "github.com/mYahagi/go-api-sample/app/domain/repository"
	"github.com/mYahagi/go-api-sample/app/infrastructure/model"
)

type BasicFareRepository struct{}

func (r BasicFareRepository) Find(endStation string) (*model.BasicFare, error) {
	var basicFare *model.BasicFare
	if endStation == "新大阪駅" {
		basicFare = &model.BasicFare{
			Name:      endStation,
			BasicFare: 8910,
		}
	} else if endStation == "姫路駅" {
		basicFare = &model.BasicFare{
			Name:      endStation,
			BasicFare: 10010,
		}
	}
	return basicFare, nil
}

func NewBasicFareRepository() iRep.IBasicFareRepository {
	return BasicFareRepository{}
}
