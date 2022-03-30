package post

import (
	"github.com/dogukanoksuz/go-rest-api-example/app/models"
	"github.com/dogukanoksuz/go-rest-api-example/platform/database"
	"github.com/gofiber/fiber/v2"
)

func Update(ctx *fiber.Ctx) error {
	request := &models.Post{}

	if err := ctx.BodyParser(&request); err != nil {
		return err
	}

	post := &models.Post{}

	err := database.Connection().First(&post, "id = ?", ctx.Params("id")).Error
	if err != nil {
		return err
	}

	err = database.Connection().Model(&post).Updates(&models.Post{
		Title:   request.Title,
		Content: request.Content,
	}).Error

	if err != nil {
		return err
	}

	return ctx.JSON(post)
}
