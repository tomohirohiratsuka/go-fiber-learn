# ベースステージ
FROM golang:1.22-alpine AS base

RUN apk add --no-cache git ca-certificates

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# 開発ステージ
FROM base AS development

# airのインストール
RUN go install github.com/cosmtrek/air@latest


CMD ["air", "-c", "docker/go/.air.toml"]

# ステージングビルドステージ
FROM base AS staging-builder

# アプリケーションのビルド（デバッグ情報を保持）
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# ステージングステージ
FROM alpine AS staging

COPY --from=staging-builder /app/main /app/main

WORKDIR /app

CMD ["./main"]

# 本番ビルドステージ
FROM base AS production-builder

# アプリケーションのビルド（最適化を施してサイズを小さくする）
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -a -installsuffix cgo -o main ./cmd/main.go

# 本番ステージ
FROM alpine AS production

RUN apk --no-cache add ca-certificates

COPY --from=production-builder /app/main /app/main

WORKDIR /app

CMD ["./main"]
