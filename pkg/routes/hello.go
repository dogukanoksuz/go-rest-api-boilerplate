package routes

import (
	"github.com/dogukanoksuz/go-rest-api-boilerplate/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func HelloRoutes(app *fiber.App) {
	app.Get("/", controllers.Hello)
}
