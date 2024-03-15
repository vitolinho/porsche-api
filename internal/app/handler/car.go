package handler

import "github.com/gofiber/fiber/v2"

func HelloCar(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{ "message": "Hello Car !", "data": nil })
}
