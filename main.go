package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gofiber/fiber"
)

func main() {

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	app.Listen(":8000")
}
