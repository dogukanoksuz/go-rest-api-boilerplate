package post

import (
	"github.com/dogukanoksuz/go-rest-api-boilerplate/app/models"
	"github.com/gofiber/fiber/v2"
)

func (p *PostController) Update(ctx *fiber.Ctx) error {
	request := &models.Post{}

	if err := ctx.BodyParser(&request); err != nil {
		return err
	}

	post := &models.Post{}

	err := p.DB.First(&post, "id = ?", ctx.Params("id")).Error
	if err != nil {
		return err
	}

	err = p.DB.Model(&post).Updates(&models.Post{
		Title:   request.Title,
		Content: request.Content,
	}).Error

	if err != nil {
		return err
	}

	return ctx.JSON(post)
}
