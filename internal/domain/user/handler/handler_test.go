package handler_test

import (
	"backend/internal/domain/user/dto"
	"backend/internal/domain/user/handler"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
	"net/http/httptest"
	"testing"
	"time"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUser(ctx context.Context, user dto.UserCreateServiceInput) (dto.UserCreateServiceOutput, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockUserService) FindUser(cts context.Context, id string) (dto.UserFindServiceOutput, error) {
	args := m.Called(id)
	return args.Get(0).(dto.UserFindServiceOutput), args.Error(1)
}

func TestUserHandler_FindUser(t *testing.T) {
	mockService := new(MockUserService)
	userId := "abcd"
	mockService.On("FindUser", userId).Return(dto.UserFindServiceOutput{
		BaseUser: dto.BaseUser{
			ID:        "id",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Email: "example@example.com",
	}, nil)
	userHandler := handler.NewUserHandler(mockService)
	app := fiber.New()
	app.Get("/user/:id", userHandler.FindUser)
	req := httptest.NewRequest("GET", fmt.Sprintf("/user/%s", userId), nil)
	resp, _ := app.Test(req)
	if resp.StatusCode != 200 {
		t.Errorf("expected status code 200, but got %d", resp.StatusCode)
	}
	mockService.AssertExpectations(t)
}
