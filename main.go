package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sanoyo/all-for-okan-go/src/database"
	"github.com/sanoyo/all-for-okan-go/src/routes"
)

func main() {
	database.Connect()
	database.AutoMigrate()

	app := fiber.New()
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	msg := "hello"
	// 	return c.SendString(msg)
	// })
	routes.Setup(app)

	app.Listen(":8000")
}
