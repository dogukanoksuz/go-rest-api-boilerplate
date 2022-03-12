package post

import (
	"github.com/dogukanoksuz/go-rest-api-example/app/models"
	"github.com/dogukanoksuz/go-rest-api-example/platform/database"
	"github.com/gofiber/fiber/v2"
)

func Show(ctx *fiber.Ctx) error {
	post := &models.Post{}

	err := database.Conn.First(&post, "id = ?", ctx.Params("id")).Error
	if err != nil {
		return err
	}

	return ctx.JSON(post)
}
