package routes

import (
	"github.com/dogukanoksuz/go-rest-api-boilerplate/app/controllers/post"
	"github.com/dogukanoksuz/go-rest-api-boilerplate/internal/database"
	"github.com/gofiber/fiber/v2"
)

func PostRoutes(app *fiber.App) {
	controller := &post.PostController{
		DB: database.Connection(),
	}

	// Post group
	postGroup := app.Group("/posts")

	// List All Posts
	postGroup.Get("/", controller.Index)

	// Show Post
	postGroup.Get("/:id", controller.Show)

	// Create Post
	postGroup.Post("/", controller.Create)

	// Update Post
	postGroup.Post("/:id", controller.Update)

	// Delete Post
	postGroup.Delete("/:id", controller.Delete)
}
