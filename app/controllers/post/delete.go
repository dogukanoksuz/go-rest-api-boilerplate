package post

import (
	"github.com/dogukanoksuz/go-rest-api-boilerplate/app/models"
	"github.com/gofiber/fiber/v2"
)

func (p *PostController) Delete(ctx *fiber.Ctx) error {
	post := &models.Post{}

	err := p.DB.First(&post, "id = ?", ctx.Params("id")).Error
	if err != nil {
		return err
	}

	err = p.DB.Delete(&post).Error
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"message": "Post deleted successfully",
	})
}
