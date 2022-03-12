package middleware

import (
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Logger(app *fiber.App) *os.File {
	basepath, _ := filepath.Abs("./")

	file, err := os.OpenFile(basepath+"/logs/main.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	app.Use(logger.New(logger.Config{
		Output:     file,
		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Europe/Istanbul",
	}))

	return file
}
