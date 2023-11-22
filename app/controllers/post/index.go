package post

import (
	"github.com/dogukanoksuz/go-rest-api-boilerplate/app/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type PostController struct {
	DB *gorm.DB
}

func (p *PostController) Index(ctx *fiber.Ctx) error {
	posts := []models.Post{}
	p.DB.Find(&posts)

	return ctx.JSON(posts)
}
