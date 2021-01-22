package controller

import (
	"fmt"
	"net/http"

	"github.com/mYahagi/go-api-sample/app/domain/usecase"
)

type PriceController struct{}

func (ctrl PriceController) Show(res http.ResponseWriter, req *http.Request) {
	start := "東京駅"
	end := req.URL.Query().Get("end")

	u := usecase.NewGetPrice(start, end)
	price, err := u.Execute()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, "error")
	}

	fmt.Fprint(res, price)
}

func NewPriceController() *PriceController {
	return &PriceController{}
}
