package post

import (
	"github.com/dogukanoksuz/go-rest-api-boilerplate/app/models"
	"github.com/dogukanoksuz/go-rest-api-boilerplate/platform/database"
	"github.com/gofiber/fiber/v2"
)

func Index(ctx *fiber.Ctx) error {
	posts := []models.Post{}
	database.Connection().Find(&posts)

	return ctx.JSON(posts)
}
