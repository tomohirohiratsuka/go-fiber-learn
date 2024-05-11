package repository

import (
	"backend/internal/domain/user/dto"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//todo test

type UserRepository interface {
	FindByID(ctx context.Context, id string) (dto.UserFindRepositoryOutput, error)
	Create(ctx context.Context, user dto.UserCreateRepositoryInput) (dto.UserCreateRepositoryOutput, error)
}

type userRepositoryImpl struct {
	user *mongo.Collection
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepositoryImpl{
		user: db.Collection("users"),
	}
}

// FindByID finds a user by ID
func (r *userRepositoryImpl) FindByID(ctx context.Context, id string) (dto.UserFindRepositoryOutput, error) {
	var user dto.UserFindRepositoryOutput

	// Create a filter to match the user by id
	filter := bson.M{"_id": id}

	// Find one user matching the filter
	result := r.user.FindOne(ctx, filter)
	if err := result.Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			// Handle the case where no documents match the filter
			return user, nil // or return an error indicating not found
		}
		// Handle other errors
		return user, err
	}

	// Decode the result into the UserFindRepositoryOutput struct
	if err := result.Decode(&user); err != nil {
		// Handle decoding error
		return user, err
	}

	return user, nil
}

// Create creates a new user
func (r *userRepositoryImpl) Create(ctx context.Context, user dto.UserCreateRepositoryInput) (dto.UserCreateRepositoryOutput, error) {
	// ドキュメントを挿入する前に、入力データからドキュメントを構築
	document := bson.M{
		"email":    user.Email,
		"password": user.Password, // 本来はパスワードはハッシュ化して保存すべき
	}

	// InsertOneメソッドを使用して新しいユーザードキュメントをコレクションに挿入
	result, err := r.user.InsertOne(ctx, document)
	if err != nil {
		// 挿入エラーを処理
		return dto.UserCreateRepositoryOutput{}, err
	}

	// InsertOneの結果から生成されたIDを取得
	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		// IDの型変換に失敗した場合のエラーハンドリング
		return dto.UserCreateRepositoryOutput{}, errors.New("failed to convert inserted ID to ObjectID")
	}

	// 成功した場合は、作成されたユーザーのID（とその他の情報があれば）を含むレスポンスオブジェクトを返す
	return dto.UserCreateRepositoryOutput{
		ID:    id.Hex(), // ObjectIDを16進数表記の文字列に変換
		Email: user.Email,
	}, nil
}
