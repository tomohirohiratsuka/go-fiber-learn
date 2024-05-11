package main

import (
	v1 "backend/api/v1"
	"backend/internal/domain/user"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func NewFiberApp() *fiber.App {
	app := fiber.New()
	// Logger Middlewareを追加
	app.Use(logger.New(logger.Config{
		// Loggerの設定をカスタマイズ
		Format:     "[${time}] ${ip}:${port} ${status} ${latency} - ${method} ${path} ${error}\n",
		TimeFormat: "2006-Jan-01 15:04:05",
		TimeZone:   "Local",
	}))

	return app
}

func AppStart(app *fiber.App) {
	go func() {
		if err := app.Listen(":8080"); err != nil {
			log.Println("Failed to start Fiber app:", err)
			os.Exit(1)
		}
	}()
}

func AppStop(app *fiber.App) {
	if err := app.Shutdown(); err != nil {
		log.Println("Failed to gracefully shutdown Fiber app:", err)
	}
}

// NewMongoDatabase は新しいmongo.Databaseインスタンスを作成して返します。
func NewMongoDatabase() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	mongoUser := os.Getenv("MONGO_USER")
	password := os.Getenv("MONGO_PASSWORD")
	port := os.Getenv("MONGO_PORT")
	host := os.Getenv("MONGO_HOST")
	database := os.Getenv("MONGO_DATABASE")
	uri := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s/%s",
		mongoUser,
		password,
		host,
		port,
		database,
	)
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 接続の確認
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}
	return client.Database(database)
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	fx.New(
		fx.Provide(
			NewFiberApp,
			NewMongoDatabase,
			v1.DefineGroup,
		),
		user.Module,
		fx.Invoke(func(lifecycle fx.Lifecycle, app *fiber.App) {
			lifecycle.Append(fx.Hook{
				OnStart: func(context.Context) error {
					AppStart(app)
					return nil
				},
				OnStop: func(ctx context.Context) error {
					AppStop(app)
					return nil
				},
			})
		}),
	).Run()
}
