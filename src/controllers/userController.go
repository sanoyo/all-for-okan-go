package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sanoyo/all-for-okan-go/src/database"
	"github.com/sanoyo/all-for-okan-go/src/models"
)

func Ambassador(c *fiber.Ctx) error {
	var users []models.User
	database.DB.Where("is_ambassador = true").Find(&users)

	return c.JSON(users)
}
