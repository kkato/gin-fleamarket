# gin-fleamarket

Ginフレームワークを使用したフリーマーケットの商品管理APIです。[Udemy Gin入門 Go言語ではじめるサーバーサイド開発](https://www.udemy.com/course/gin-golang/)の一環として作成しました。

## プロジェクト概要

このプロジェクトは、Go言語のGinフレームワークを使用したRESTful APIアプリケーションです。フリーマーケットの商品管理とユーザー認証機能を提供します。

## 機能

- **商品管理**
  - 商品の一覧表示（認証不要）
  - 商品の詳細表示（認証必要）
  - 商品の作成・更新・削除（認証必要）
- **ユーザー認証**
  - ユーザー登録（サインアップ）
  - ユーザーログイン
  - JWT認証
- **セキュリティ**
  - パスワードのハッシュ化
  - JWTトークンによる認証ミドルウェア
  - CORS対応

## 技術スタック

- **言語**: Go 1.24.0
- **Webフレームワーク**: Gin
- **ORM**: GORM
- **データベース**: PostgreSQL
- **認証**: JWT (golang-jwt/jwt)
- **パスワードハッシュ**: bcrypt
- **コンテナ**: Docker / Docker Compose

## アーキテクチャ

レイヤードアーキテクチャを採用：

- **Controllers**: HTTPリクエストの処理
- **Services**: ビジネスロジック
- **Repositories**: データアクセス層
- **Models**: データモデル定義
- **DTOs**: データ転送オブジェクト
- **Middlewares**: 認証などの共通処理

## セットアップ手順

### リポジトリのクローン

```bash
git clone https://github.com/kkato/gin-fleamarket.git
cd gin-fleamarket
```

### 環境変数の設定

`.env`ファイルを作成し、以下の環境変数を設定：

```bash
ENV=dev
DB_HOST=localhost
DB_PORT=5432
DB_USER=ginuser
DB_PASSWORD=ginpassword
DB_NAME=fleamarket
SECRET_KEY=your-secret-key-here
```

### Docker Composeで起動

PostgreSQLとpgAdminをDocker Composeで起動：

```bash
docker-compose up -d
```

これにより以下のサービスが起動します：
- PostgreSQL: `localhost:5432`
- pgAdmin: `http://localhost:81`
  - Email: gin@example.com
  - Password: ginpassword

### 依存関係のインストール

```bash
go mod tidy
```

### アプリケーションの起動

```bash
go run main.go
```

アプリケーションは `http://localhost:8080` で起動します。

### 開発モード（ホットリロード）

Air を使用してホットリロード機能を有効にする場合：

```bash
air
```

## API エンドポイント

### 認証

- `POST /auth/signup` - ユーザー登録
- `POST /auth/login` - ログイン

### 商品

- `GET /items` - 商品一覧取得（認証不要）
- `GET /items/:id` - 商品詳細取得（認証必要）
- `POST /items` - 商品作成（認証必要）
- `PUT /items/:id` - 商品更新（認証必要）
- `DELETE /items/:id` - 商品削除（認証必要）

認証が必要なエンドポイントには、リクエストヘッダーに以下を含める必要があります：

```
Authorization: Bearer <JWT_TOKEN>
```

## テスト

テストの実行：

```bash
go test -v ./...
```

## プロジェクト構成

```
.
├── controllers/       # HTTPハンドラ
├── services/         # ビジネスロジック
├── repositories/     # データアクセス層
├── models/          # データモデル
├── dto/             # データ転送オブジェクト
├── middlewares/     # ミドルウェア
├── infra/           # インフラストラクチャ（DB接続など）
├── migrations/      # データベースマイグレーション
├── docker/          # Docker設定ファイル
├── main.go          # エントリーポイント
└── docker-compose.yaml
```
