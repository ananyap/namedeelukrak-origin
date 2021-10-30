package handlers

import (
	"github.com/ananyap/namedeelukrak/services"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type numberHandler struct {
	numberSrv services.NumberService
}

func NewNumberHandler(numberSrv services.NumberService) numberHandler {
	return numberHandler{numberSrv: numberSrv}
}

func (h numberHandler) GetNumbers(c *fiber.Ctx) error {
	numbers, err := h.numberSrv.GetNumbers()
	if err != nil {
		logrus.Print("Internal Error")
	}

	c.JSON(numbers)

	return nil
}
