package post

import (
	"github.com/dogukanoksuz/go-rest-api-boilerplate/app/models"
	"github.com/dogukanoksuz/go-rest-api-boilerplate/platform/database"
	"github.com/gofiber/fiber/v2"
)

func Create(ctx *fiber.Ctx) error {
	post := models.Post{}
	if err := ctx.BodyParser(&post); err != nil {
		return err
	}

	if err := database.Connection().Create(&post).Error; err != nil {
		return err
	}

	return ctx.JSON(post)
}
