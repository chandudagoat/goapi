package main

import (
	"goapi/database"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := init_app()

	if err != nil {
		panic(err)
	}

	// defer close mongo db
	defer database.CloseMongo()

	app := fiber.New()
	groups(app)
	routes(app)

	port := os.Getenv("PORT")
	app.Listen(":" + port)
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
	go_env := os.Getenv("GO_ENV")
	if go_env == "" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}
	return nil
}
