package controller

import (
	"fmt"
	"net/http"

	"github.com/mYahagi/go-api-sample/app/domain/usecase"
)

type PriceController struct{}

func (ctrl PriceController) Show(res http.ResponseWriter, req *http.Request) {
	// TODO bindingした値をstring型の引数に渡すと死ぬのでちゃんとした方法探してくる
	//var endStation string
	//endStation = req.FormValue("end")

	u := usecase.NewGetPrice("大阪駅")
	price, err := u.Execute()
	if err != nil {
		fmt.Fprint(res, "error")
	}

	fmt.Fprint(res, price)
}

func NewPriceController() *PriceController {
	return &PriceController{}
}
