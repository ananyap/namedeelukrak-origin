package handlers

import (
	"strconv"

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
func (h numberHandler) GetNumber(c *fiber.Ctx) error {
	idx := c.Params("id")
	numberId, _ := strconv.Atoi(idx)

	number, err := h.numberSrv.GetNumber(numberId)
	if err != nil {
		logrus.Print("Internal Error")
	}

	c.JSON(number)

	return nil
}
