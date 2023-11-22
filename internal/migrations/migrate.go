package migrations

import (
	"github.com/dogukanoksuz/go-rest-api-boilerplate/app/models"
	"github.com/dogukanoksuz/go-rest-api-boilerplate/internal/database"
)

func Migrate() {
	database.Connection().AutoMigrate(&models.Post{})
}
