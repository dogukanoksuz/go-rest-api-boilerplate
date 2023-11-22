package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/dogukanoksuz/go-rest-api-boilerplate/app/middleware"
	"github.com/dogukanoksuz/go-rest-api-boilerplate/app/routes"
	"github.com/dogukanoksuz/go-rest-api-boilerplate/internal/migrations"
	"github.com/dogukanoksuz/go-rest-api-boilerplate/internal/server"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	if !fiber.IsChild() {
		// Application initialization
		migrations.Migrate()
	}

	port, _ := strconv.Atoi(os.Getenv("APP_PORT"))

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
