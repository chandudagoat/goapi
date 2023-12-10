package main

import (
	"goapi/database"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := init_app()

	if err != nil {
		panic(err)
	}

	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World")
	})

	app.Listen(":3000")
}

func init_app() error {
	err := load_env()
	if err != nil {
		return err
	}

	err = database.LoadMongo()
	if err != nil {
		return err
	}

	return nil
}

func load_env() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}
