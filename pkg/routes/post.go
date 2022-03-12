package routes

import (
	"github.com/dogukanoksuz/go-rest-api-example/app/controllers/post"
	"github.com/gofiber/fiber/v2"
)

func PostRoutes(app *fiber.App) {
	// List All Posts
	app.Get("/posts", post.Index)

	// Show Post
	app.Get("/post/:id", post.Show)

	// Create Post
	app.Post("/post", post.Create)

	// Update Post
	app.Post("/post/:id", post.Update)

	// Delete Post
	app.Delete("/post/:id", post.Delete)
}
