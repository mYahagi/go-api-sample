package controller

import (
	"fmt"
	"net/http"
)

type PriceController struct{}

func (ctrl PriceController) Show(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
}

func NewPriceController() *PriceController {
	return &PriceController{}
}
