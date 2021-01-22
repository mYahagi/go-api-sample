package repository

import "github.com/mYahagi/go-api-sample/app/infrastructure/model"

type IExpressChargeRepository interface {
	// TODO 適切なentity作ったらそれを返すようにする
	Find() *model.ExpressCharge
}
