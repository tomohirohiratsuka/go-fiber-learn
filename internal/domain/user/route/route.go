package route

import (
	"backend/internal/domain/user/handler"
	"github.com/gofiber/fiber/v2"
)

// RegisterUserRoutes はユーザー関連のルートを設定します。
func RegisterUserRoutes(v1Group fiber.Router, uh *handler.UserHandler) {
	user := v1Group.Group("/users")
	user.Post("/", func(c *fiber.Ctx) error {
		return uh.CreateUser(c)
	})
	user.Get("/:id", func(c *fiber.Ctx) error {
		return uh.FindUser(c)
	})
}
