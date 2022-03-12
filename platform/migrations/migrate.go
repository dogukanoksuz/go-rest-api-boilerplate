package migrations

import (
	"github.com/dogukanoksuz/go-rest-api-example/app/models"
	"github.com/dogukanoksuz/go-rest-api-example/platform/database"
)

func Migrate() {
	database.Conn.AutoMigrate(&models.Post{})
}
