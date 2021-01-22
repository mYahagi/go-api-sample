package repository

import "github.com/mYahagi/go-api-sample/app/infrastructure/model"

type IBasicFareRepository interface {
	// TODO 適切なentity作ったらそれを返すようにする
	Find(endStation string) model.BasicFare
}
