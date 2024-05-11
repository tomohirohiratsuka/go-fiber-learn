package handler

import (
	"backend/internal/domain/user/dto"
	"backend/internal/domain/user/service"
	"backend/internal/shared/util"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service.UserService
}

// NewUserHandler は UserHandler の新しいインスタンスを作成します。
func NewUserHandler(us service.UserService) *UserHandler {
	return &UserHandler{UserService: us}
}

func (r *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req dto.UserCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Request parsing failed"})
	}
	if err := req.Validate(); err != nil {
		return util.HandleValidationErrors(c, err)
	}
	var createInput dto.UserCreateServiceInput
	createInput.FromRequest(req)
	ctx := c.Context()
	var user, _ = r.UserService.CreateUser(ctx, createInput)
	return c.Status(fiber.StatusCreated).JSON(user)
}

func (r *UserHandler) FindUser(c *fiber.Ctx) error {
	ctx := c.Context()
	var user, _ = r.UserService.FindUser(ctx, c.Params("id"))
	return c.JSON(user)
}
