package controller

import (
	"fmt"
	"net/http"
)

type PriceController struct{}

func (ctrl PriceController) Show(res http.ResponseWriter, req *http.Request) {
	endStation := req.FormValue("end")
	fmt.Fprint(res, endStation)
}

func NewPriceController() *PriceController {
	return &PriceController{}
}
