package utils

import (
	"fmt"
	"log"

	"github.com/dogukanoksuz/go-rest-api-example/pkg/middleware"
	"github.com/dogukanoksuz/go-rest-api-example/pkg/routes"
	"github.com/gofiber/fiber/v2"
)

func CreateServer(port int) {
	// Create Fiber App
	app := fiber.New()

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
