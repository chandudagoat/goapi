package handlers

import (
	"context"
	"fmt"
	"goapi/database"

	"github.com/gofiber/fiber/v2"
)

type libraryDTO struct {
	Name    string `json: "name" bson: "name"`
	Address string `json: "address" bson: "address"`
}

func CreateLibrary(ctx *fiber.Ctx) error {
	new_lib := new(libraryDTO)

	if err := ctx.BodyParser(new_lib); err != nil {
		return err
	}

	fmt.Println(new_lib)
	libCollection := database.GetCollection("libraries")
	new_doc, err := libCollection.InsertOne(context.TODO(), new_lib)

	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{"id": new_doc.InsertedID})
}
