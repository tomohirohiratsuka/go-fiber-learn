package service

import (
	"backend/internal/domain/user/dto"
	"backend/internal/domain/user/repository"
	"context"
	"time"
)

//todo test

// UserService - インターフェースの定義
type UserService interface {
	CreateUser(ctx context.Context, user dto.UserCreateServiceInput) (dto.UserCreateServiceOutput, error)
	FindUser(ctx context.Context, id string) (dto.UserFindServiceOutput, error)
}

// userServiceImpl - UserServiceインターフェースの具体的な実装
type userServiceImpl struct {
	UserRepository repository.UserRepository
}

// NewUserService - UserServiceImplのインスタンスを生成するコンストラクタ関数
func NewUserService(ur repository.UserRepository) UserService {
	return &userServiceImpl{UserRepository: ur}
}

// CreateUser - Userの作成処理を行う
func (s *userServiceImpl) CreateUser(ctx context.Context, user dto.UserCreateServiceInput) (dto.UserCreateServiceOutput, error) {
	createInput := dto.UserCreateRepositoryInput{
		Email:    user.Email,
		Password: user.Password,
	}
	repositoryOutput, err := s.UserRepository.Create(ctx, createInput)
	if err != nil {
		return dto.UserCreateServiceOutput{}, err
	}
	// ここにビジネスロジックを実装
	output := dto.UserCreateServiceOutput{
		BaseUser: dto.BaseUser{
			ID:        repositoryOutput.ID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Email: "example@example.com",
	}

	return output, nil
}

// FindUser - 指定されたIDのUserを検索する
func (s *userServiceImpl) FindUser(ctx context.Context, id string) (dto.UserFindServiceOutput, error) {
	// ここにビジネスロジックを実装
	repositoryOutput, err := s.UserRepository.FindByID(ctx, id)
	if err != nil {
		return dto.UserFindServiceOutput{}, err
	}

	output := dto.UserFindServiceOutput{
		BaseUser: dto.BaseUser{
			ID:        repositoryOutput.ID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Email: repositoryOutput.Email,
	}
	return output, nil
}
