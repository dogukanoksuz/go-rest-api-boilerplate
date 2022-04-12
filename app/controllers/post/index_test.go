package post

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dogukanoksuz/go-rest-api-boilerplate/app/models"
	"github.com/gofiber/fiber/v2"
)

func TestIndex(t *testing.T) {
	type wanted struct {
		statusCode int
	}

	// Create fiber app for testing purposes
	app := fiber.New()
	app.Get("/", Index)

	// Tests struct
	tests := []struct {
		name        string
		description string
		endpoint    string
		want        wanted
	}{
		{
			name:        "Get posts list",
			description: "This test should return 200 status code and return a list of posts",
			endpoint:    "/",
			want: wanted{
				statusCode: 200,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", tt.endpoint, strings.NewReader(""))

			// Create request
			res, err := app.Test(req)
			if err != nil {
				t.Errorf("Cannot test Fiber handler: %v", err)
				t.Fail()
			}

			// Assertions
			if res.StatusCode != tt.want.statusCode {
				t.Errorf("Expected status code %d, got %d", tt.want.statusCode, res.StatusCode)
			}

			answer, err := io.ReadAll(res.Body)
			if err != nil {
				t.Errorf("Cannot parse body: %v", err)
			}

			var posts []models.Post
			if err := json.Unmarshal(answer, &posts); err != nil {
				t.Errorf("Cannot unmarshal body, might returned wrong type of output: %v", err)
			}
		})
	}
}
