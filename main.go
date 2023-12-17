package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()
	env := godotenv.Load()

	if env != nil {
		log.Fatal("couldn't load .env file")
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("nuh uh")
	})

	PORT := os.Getenv("PORT")
	app.Listen(":" + PORT)
}
