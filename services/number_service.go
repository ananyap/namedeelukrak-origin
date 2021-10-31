package services

import (
	"github.com/ananyap/namedeelukrak/repositories"
	"github.com/sirupsen/logrus"
)

type numberService struct {
	numberRepo repositories.NumberRepository
}

func NewNumberService(numberRepo repositories.NumberRepository) NumberService {
	return numberService{numberRepo}
}

func (service numberService) GetNumbers() ([]NumberResponse, error) {
	numbers, err := service.numberRepo.GetAll()
	if err != nil {
		logrus.Error(err)
	}

	numberRess := []NumberResponse{}
	for _, number := range numbers {
		numberRes := NumberResponse{
			PairNumber: number.PairNumber,
			PairType:   number.PairType,
			PairPoint:  number.PairPoint,
		}
		numberRess = append(numberRess, numberRes)
	}

	return numberRess, nil
}

func (service numberService) GetNumber(id int) (*NumberResponse, error) {
	number, err := service.numberRepo.GetById(id)
	if err != nil {
		logrus.Error(err)
	}

	numberRes := NumberResponse{
		PairNumber: number.PairNumber,
		PairType:   number.PairType,
		PairPoint:  number.PairPoint,
	}

	return &numberRes, nil
}
