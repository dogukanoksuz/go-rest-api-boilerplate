package utils

import (
	"fmt"
	"log"

	"github.com/dogukanoksuz/go-rest-api-boilerplate/pkg/middleware"
	"github.com/dogukanoksuz/go-rest-api-boilerplate/pkg/routes"
	"github.com/dogukanoksuz/go-rest-api-boilerplate/platform/server"
	"github.com/gofiber/fiber/v2"
)

func CreateServer(port int) {
	// Create Fiber App
	app := fiber.New(fiber.Config{
		Prefork:      true,
		ErrorHandler: server.ErrorHandler,
	})

	// Initialize Logger
	file := middleware.Logger(app)
	defer file.Close()

	// Middlewares
	app.Use(middleware.Example)

	// Mount routes
	routes.HelloRoutes(app)
	routes.PostRoutes(app)

	// Start server
	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
