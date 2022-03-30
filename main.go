package main

import (
	"os"
	"strconv"

	"github.com/dogukanoksuz/go-rest-api-boilerplate/pkg/utils"
	"github.com/dogukanoksuz/go-rest-api-boilerplate/platform/migrations"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	if !fiber.IsChild() {
		// Application initialization
		migrations.Migrate()
	}

	port, _ := strconv.Atoi(os.Getenv("APP_PORT"))

	utils.CreateServer(port)
}
