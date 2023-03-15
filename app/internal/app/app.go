package app

import (
	"time"

	"project/internal/setup/routes"
	"project/internal/utils/response"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sirupsen/logrus"
)

func NewApp(logger *logrus.Logger) (app *fiber.App) {

	app = fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			logger.WithFields(logrus.Fields{
				"baseUrl": ctx.Hostname() + ctx.OriginalURL(),
				"user":    ctx.Get("user"),
			}).Error(err.Error())
			res := response.Error("Fiber error handler", err.Error())
			result := ctx.Status(code).JSON(res)
			return result
		},

		AppName:                 "My Project",
		BodyLimit:               4 * 1024 * 1024,
		EnableTrustedProxyCheck: true,
		ServerHeader:            "Döwlet hyzmatlar portaly",
		WriteTimeout:            time.Second * 40,
		ReadTimeout:             time.Second * 40,
	})

	// app.Use(cors.New())

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,API-KEY",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	routes.SetAllAdminRoutes(app)

	return app
}
