package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func Example(ctx *fiber.Ctx) error {
	log.Println("My middleware is working!")

	return ctx.Next()
}
