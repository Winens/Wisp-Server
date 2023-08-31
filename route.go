package main

import (
	"github.com/Winens/Wisp-Server/handler"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	rAuth := app.Group("/auth")
	{
		rAuth.Post("/signup", handler.SignUp)
	}
}
