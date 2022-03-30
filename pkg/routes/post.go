package routes

import (
	"github.com/dogukanoksuz/go-rest-api-example/app/controllers/post"
	"github.com/gofiber/fiber/v2"
)

func PostRoutes(app *fiber.App) {
	// Post group
	postGroup := app.Group("/posts")

	// List All Posts
	postGroup.Get("/", post.Index)

	// Show Post
	postGroup.Get("/:id", post.Show)

	// Create Post
	postGroup.Post("/", post.Create)

	// Update Post
	postGroup.Post("/:id", post.Update)

	// Delete Post
	postGroup.Delete("/:id", post.Delete)
}
