package handlers

import (
	"context"
	"fmt"
	"goapi/database"
	"goapi/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
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

func GetLibraries(ctx *fiber.Ctx) error {
	libCollection := database.GetCollection("libraries")
	cursor, err := libCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		return err
	}

	var libraries models.Library

	if err = cursor.All(context.TODO(), &libraries); err != nil {
		return err
	}

	return ctx.JSON(libraries)
}
