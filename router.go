package main

import (
	"goapi/handlers"

	"github.com/gofiber/fiber/v2"
)

func routes(app *fiber.App) error {
	// create a health check route
	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Ok")
	})

	return nil
}

func groups(app *fiber.App) error {
	// create the library routes
	libGroup := app.Group("/library")
	libGroup.Get("/test", handlers.TestHandler)

	return nil
}
