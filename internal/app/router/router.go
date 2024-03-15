package router

import (
	"porsche-api/internal/app/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func MakeRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	v1 := api.Group("/v1")
	v1.Get("/cars", handler.HelloCar)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})
}
