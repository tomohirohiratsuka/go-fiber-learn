package util

import (
	"backend/internal/shared/dto"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// HandleValidationErrors はバリデーションエラーを受け取り、
// 適切なフォーマットでエラーレスポンスを返す関数です。
func HandleValidationErrors(c *fiber.Ctx, err error) error {
	var errors []*dto.ValidationErrorResponse
	for _, err := range err.(validator.ValidationErrors) {
		var element dto.ValidationErrorResponse
		element.FailedField = err.StructNamespace()
		element.Tag = err.Tag()
		element.Value = fmt.Sprintf("%v", err.Value())

		errors = append(errors, &element)
	}
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errors})
}
