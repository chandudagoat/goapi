package main

import (
	"goapi/handlers"

	"github.com/gofiber/fiber/v2"
)

func routes(app *fiber.App) {
	// create a health check route
	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Ok")
	})
}

func groups(app *fiber.App) {
	// create the library routes
	libGroup := app.Group("/library")

	// get reqs
	libGroup.Get("/test", handlers.TestHandler)
	libGroup.Get("/get_libs", handlers.GetLibraries)

	// post req
	libGroup.Post("/create", handlers.CreateLibrary)
}
