FROM golang:latest

# 作業ディレクトリの設定
WORKDIR /app

# 依存関係ファイルのコピー
COPY go.mod go.sum ./

# 依存関係のダウンロード
RUN go mod download

# ソースコードのコピー
COPY . .

# サーバーのビルド
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

# サーバーアプリケーションを起動
ENTRYPOINT ["./server"]
