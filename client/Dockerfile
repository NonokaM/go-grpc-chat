FROM golang:latest

# 作業ディレクトリの設定
WORKDIR /app

# 依存関係ファイルのコピー
COPY go.mod go.sum ./

# 依存関係のダウンロード
RUN go mod download

# ソースコードのコピー
COPY . .

# クライアントのビルド
RUN CGO_ENABLED=0 GOOS=linux go build -o client .

# 実行ファイルに実行権限を与える
RUN chmod +x client

# クライアントアプリケーションを対話型で起動
ENTRYPOINT ["./client"]
