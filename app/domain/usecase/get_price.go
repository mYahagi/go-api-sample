package usecase

type GetPrice struct {
	endStation string
}

func (u GetPrice) Execute() (int, error) {
	return 2000, nil
}

func NewGetPrice(endStation string) *GetPrice {
	return &GetPrice{
		endStation: endStation,
	}
}
