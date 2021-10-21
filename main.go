package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	_, err := gorm.Open(mysql.Open("root:root@tcp(db:3306)/ambassador"), &gorm.Config{})
	if err != nil {
		panic("Could not connect with the database!")
	}

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		msg := "hello"
		return c.SendString(msg)
	})

	app.Listen(":8000")
}
