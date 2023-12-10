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

	// defer close mongo db
	defer database.CloseMongo()

	app := fiber.New()
	groups(app)
	routes(app)

	// app.Post("/", func(ctx *fiber.Ctx) error {
	// 	sampleDoc := bson.M{"name": "hello there"}
	// 	collection := database.GetCollection("todos")
	// 	newDoc, err := collection.InsertOne(context.TODO(), sampleDoc)

	// 	if err != nil {
	// 		return ctx.Status(fiber.StatusInternalServerError).SendString("error inserting todo")
	// 	}
	// 	return ctx.JSON(newDoc)
	// })

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
