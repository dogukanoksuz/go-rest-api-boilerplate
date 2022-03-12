package main

import (
	"os"
	"strconv"

	"github.com/dogukanoksuz/go-rest-api-example/pkg/utils"
	"github.com/dogukanoksuz/go-rest-api-example/platform/database"
	"github.com/dogukanoksuz/go-rest-api-example/platform/migrations"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// Application initialization
	database.Init()
	migrations.Migrate()

	port, _ := strconv.Atoi(os.Getenv("APP_PORT"))

	utils.CreateServer(port)
}
