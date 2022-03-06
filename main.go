package main

import (
	"log"

	"github.com/dogukanoksuz/go-rest-api-example/app/controllers"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := fiber.New()

	app.Get("/", controllers.Hello)

	log.Fatal(app.Listen(":3000"))
}
