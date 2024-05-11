package v1

import (
	"github.com/gofiber/fiber/v2"
)

func DefineGroup(app *fiber.App) fiber.Router {
	return app.Group("/api/v1")
}
