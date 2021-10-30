package services

type NumberResponse struct {
	PairNumber string `json:"pairnumber"`
	PairType   string `json:"pairtype"`
	PairPoint  int    `json:"pairpoint"`
}

type NumberService interface {
	GetNumbers() ([]NumberResponse, error)
	GetNumber() (*NumberResponse, error)
}
