package routes

import (
	xmlConstructor "project/internal/domain/xml/constructor"

	"github.com/gofiber/fiber/v2"
)

func SetAllAdminRoutes(app *fiber.App) {
	xml := app.Group("/xml")
	xml.Get("/import", xmlConstructor.XmlController.Import)
	xml.Get("/all", xmlConstructor.XmlController.All)
}
