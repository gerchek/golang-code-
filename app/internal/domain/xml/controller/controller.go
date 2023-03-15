package controller

import (
	"net/http"
	"project/internal/domain/xml/service"
	"project/internal/utils/response"

	"github.com/gofiber/fiber/v2"
)

type XmlController interface {
	All(ctx *fiber.Ctx) error
	Import(ctx *fiber.Ctx) error
}

type xmlController struct {
	service service.XmlService
}

func NewXmlController(service service.XmlService) XmlController {
	return &xmlController{
		service: service,
	}
}

// GET ALL
func (c *xmlController) All(ctx *fiber.Ctx) error {

	result := c.service.All()
	return ctx.Status(http.StatusOK).JSON(result)

}

// IMPORT
func (c *xmlController) Import(ctx *fiber.Ctx) error {

	err := c.service.Import()

	if err != nil {
		res := response.Error("An error has occurred", err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}
	res := response.Success(true, "Success")
	return ctx.Status(http.StatusOK).JSON(res)

}
