package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sanoyo/all-for-okan-go/src/database"
	"github.com/sanoyo/all-for-okan-go/src/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "passwprd is mismatched",
		})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 12)
	user := models.User{
		FirstName:    data["first_name"],
		LastName:     data["last_name"],
		Email:        data["email"],
		Password:     password,
		IsAmbassador: false,
	}

	database.DB.Create(&user)

	return c.JSON(user)
}
