package utils

import (
	"fmt"
	"log"

	"github.com/dogukanoksuz/go-rest-api-example/pkg/routes"
	"github.com/gofiber/fiber/v2"
)

func CreateServer(port int) {
	app := fiber.New()

	routes.HelloRoutes(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
