package main

import (
	"porsche-api/internal/infrastructure/database"
	"porsche-api/internal/app/router"
	"github.com/gofiber/fiber/v2"
	"log"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Database()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000/",
		AllowHeaders:     "Origin, Content-Type, Accept, Content-Length, Accept-Language, Cache-Control, Accept-Encoding, Connection, Access-Control-Allow-Origin, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowCredentials: true,
	}))

	router.MakeRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
