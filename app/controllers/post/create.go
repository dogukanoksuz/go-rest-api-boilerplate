package post

import (
	"github.com/dogukanoksuz/go-rest-api-example/app/models"
	"github.com/dogukanoksuz/go-rest-api-example/platform/database"
	"github.com/gofiber/fiber/v2"
)

func Create(ctx *fiber.Ctx) error {
	post := models.Post{}
	if err := ctx.BodyParser(&post); err != nil {
		return ctx.Status(503).JSON(err)
	}

	if err := database.Connection().Create(&post).Error; err != nil {
		return ctx.Status(503).JSON(err)
	}

	return ctx.JSON(post)
}
