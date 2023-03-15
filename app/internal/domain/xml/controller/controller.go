package controller

import (
	"fmt"
	"net/http"
	"project/internal/domain/xml/dto"
	"project/internal/domain/xml/service"
	"project/internal/utils/customvalidator"
	"project/internal/utils/response"

	"github.com/gofiber/fiber/v2"
)

type XmlController interface {
	All(ctx *fiber.Ctx) error
	Import(ctx *fiber.Ctx) error
	GetNames(ctx *fiber.Ctx) error
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
	res := response.Success(true, "Success", nil)
	return ctx.Status(http.StatusOK).JSON(res)

}

// GetNames
func (c *xmlController) GetNames(ctx *fiber.Ctx) error {

	var xmlDTO dto.XmlDTO
	err := ctx.BodyParser(&xmlDTO)
	fmt.Println(&xmlDTO)
	if err != nil {
		res := response.Error("Bad request", err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}
	err = customvalidator.ValidateStruct(&xmlDTO)
	if err != nil {
		res := response.Error("Validation error", err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}
	data, err := c.service.GetNames(&xmlDTO)
	if err != nil {
		res := response.Error("An error has occurred", err.Error())
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}
	res := response.Success(true, "Success", data)
	return ctx.Status(http.StatusOK).JSON(res)
}
