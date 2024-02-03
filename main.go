package main

import (
	"github.com/gofiber/fiber/v2"
	"main.go/routers"
)

func main() {
  app := fiber.New()

  routers.RouterApp(app)
  app.Listen(":8000")
}

