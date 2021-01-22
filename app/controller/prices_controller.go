package controller

import (
	"fmt"
	"net/http"

	usecase2 "github.com/mYahagi/go-api-sample/app/domain/usecase"
)

type PriceController struct{}

func (ctrl PriceController) Show(res http.ResponseWriter, req *http.Request) {
	endStation := req.FormValue("end")
	usecase := usecase2.NewGetPrice(endStation)
	price, err := usecase.Execute()

	if err != nil {
		fmt.Fprint(res, "error")
	}
	fmt.Fprint(res, price)
}

func NewPriceController() *PriceController {
	return &PriceController{}
}
