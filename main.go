package main

import (
	"os"
	"strconv"

	"github.com/dogukanoksuz/go-rest-api-example/pkg/utils"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	port, _ := strconv.Atoi(os.Getenv("APP_PORT"))

	utils.CreateServer(port)
}
