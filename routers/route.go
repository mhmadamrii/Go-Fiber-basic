package routers

import (
	"github.com/gofiber/fiber/v2"
	"main.go/controllers"
)

func RouterApp(c *fiber.App) {
	c.Get("/", controllers.UserControllerShow)
}
