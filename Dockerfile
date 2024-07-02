# Build the Go app
FROM golang:1.22.4 as builder
WORKDIR /app
# 依存関係を含むモジュール情報を/app ディレクトリにコピーし、コンテナ内に取り込む
COPY src/go.mod src/go.sum ./
# 依存するモジュールをダウンロード
RUN go mod download
# アプリケーションのソースコードがコンテナ内に配置
COPY src/ ./
# アプリケーションをビルド
RUN go build -o /todo-app

# Deploy the Go app
FROM golang:1.22.4
WORKDIR /
# ビルドされたアプリケーションのバイナリが実行用のコンテナ内に配置される
COPY --from=builder /todo-app /todo-app
# Dockerコンテナがポート8080でリッスンすることを宣言
EXPOSE 8080
# コンテナが起動した際に実行されるデフォルトのコマンドを指定。/todo-appコマンドよりコンテナが起動し、アプリケーションが実行される
CMD ["/todo-app"]
