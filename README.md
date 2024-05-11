# Restful API Template

# How to use
## 1. Clone this repository
```bash
git clone
```
## 2. Copy .env.example to .env
```bash
cp .env.example .env
```

## 3. Run the server
```bash
docker-compose up
```

## Overview

このドキュメントでは、プロジェクトの基本構造と、パッケージ名やエイリアス名に関するガイドラインについて説明します。

## Project Structure

プロジェクトは以下の基本構造に従っています:

```
.
├── api
│   └── v1
│       └── user.go
├── cmd
│   └── main.go
├── go.mod
├── go.sum
└── internal
    ├── config
    ├── domain
    │   └── user
    │       ├── dto
    │       │   ├── request.go
    │       │   └── response.go
    │       ├── handler
    │       │   └── handler.go
    │       ├── repository
    │       │   └── repository.go
    │       └── service
    │           └── service.go
    └── shared
        └── dto
            └── response.go
```


## Package and Alias Naming Guidelines

### パッケージ名

- パッケージ名は、そのパッケージの提供する機能や目的を簡潔に反映するものにします。
- 一般的に認識されている名前（例: `fmt`, `http`）や、短く明瞭な名前（例: `user`, `service`）をそのまま使用します。
- パッケージ名は基本的には単数形を使用し、複数形は避けます。

### エイリアス名

- エイリアスは、名前の衝突を避ける、またはコードの可読性を向上させるためにのみ使用します。
- エイリアスを使用する場合は、そのエイリアスが何を指しているのかを一目で理解できるような明確で記述的な名前を選びます。
- 使用されるコンテキストにおいてサブとなるドメインのサービスやDTOにエイリアスを付与します。主ドメインはエイリアスなしで利用し、他のサブドメインに対しては区別のためにエイリアスを適用します。

## DTO Naming Convention
### 命名規則の構成要素
 - ドメイン名 (Domain): 対象となるドメイン（例：User、Orderなど）。 
 - 操作 (Operation): 実行される操作（例：Create、Update、Deleteなど）。 
 - I/O: 入力か出力かを示す（例：Input、Output）。 
 - レイヤー (Layer): 使用されるレイヤーを示す（例：Service、Repository）。

### ベースフォーマット
```
[Domain][Operation][Layer][I/O]
```
### 具体的な例
サービス層の入力DTO:
 - UserCreateServiceInput: ユーザー作成のためのサービス層の入力DTO。 
 - OrderUpdateServiceInput: 注文更新のためのサービス層の入力DTO。

サービス層の出力DTO:
 - UserGetServiceOutput: ユーザー取得のためのサービス層の出力DTO。 
 - OrderDeleteServiceOutput: 注文削除のためのサービス層の出力DTO。

リポジトリ層の入力DTO:
 - UserCreateRepositoryInput: ユーザー作成のためのリポジトリ層の入力DTO。 
 - OrderUpdateRepositoryInput: 注文更新のためのリポジトリ層の入力DTO。

リポジトリ層の出力DTO:
 - UserGetRepositoryOutput: ユーザー取得のためのリポジトリ層の出力DTO。 
 - OrderDeleteRepositoryOutput: 注文削除のためのリポジトリ層の出力DTO。

規定のDTO:
 - BaseUser: ユーザーの基本DTO。

## Linter
 - [golangci-lint](https://golangci-lint.run/)

## References
 - [Effective Go](https://go.dev/doc/effective_go)
 - [Go Style Best Practice](https://google.github.io/styleguide/go/best-practices.html)
 - [Twelve Go Best Practices](https://go.dev/talks/2013/bestpractices.slide#1)