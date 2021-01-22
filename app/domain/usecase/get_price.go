package usecase

type GetPrice struct {
	startStation string
	endStation   string
}

func (u GetPrice) Execute() (int, error) {
	if u.endStation == "大阪駅" {
		return 3000, nil
	}
	return 2000, nil
}

func NewGetPrice(startStation, endStation string) *GetPrice {
	return &GetPrice{
		startStation: startStation,
		endStation:   endStation,
	}
}
