package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sanoyo/all-for-okan-go/src/controllers"
	"github.com/sanoyo/all-for-okan-go/src/middleware"
)

func Setup(app *fiber.App) {
	api := app.Group("api")
	admin := api.Group("admin")
	admin.Post("register", controllers.Register)
	admin.Post("login", controllers.Login)

	adminAuthenticated := admin.Use(middleware.IsAuthenticated)
	adminAuthenticated.Get("user", controllers.User)
	adminAuthenticated.Post("logout", controllers.Logout)
	adminAuthenticated.Put("user/info", controllers.UpdateProfile)
	adminAuthenticated.Put("user/password", controllers.UpdatePasword)
}
