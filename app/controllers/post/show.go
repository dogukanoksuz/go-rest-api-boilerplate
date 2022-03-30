package post

import (
	"github.com/dogukanoksuz/go-rest-api-boilerplate/app/models"
	"github.com/dogukanoksuz/go-rest-api-boilerplate/platform/database"
	"github.com/gofiber/fiber/v2"
)

func Show(ctx *fiber.Ctx) error {
	post := &models.Post{}

	err := database.Connection().First(&post, "id = ?", ctx.Params("id")).Error
	if err != nil {
		return err
	}

	return ctx.JSON(post)
}
