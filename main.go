package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sanoyo/all-for-okan-go/src/database"
)

func main() {
	database.Connect()
	database.AutoMigrate()

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		msg := "hello"
		return c.SendString(msg)
	})

	app.Listen(":8000")
}
