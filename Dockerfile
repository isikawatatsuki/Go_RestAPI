# Go 1.22のベースイメージを使用
# 注意すること go.mod記載のversionと合わせる
# alpineを使用する(軽量化を目的としている)

FROM golang:1.24-alpine

# コンテナ内の作業ディレクトリを/appに設定する
WORKDIR /app

# ホストマシンからtodo-app内の全ファイルを、コンテナ内の作業ディレクトリ(/app)にコピー
COPY .  ./

# main.goファイルをコンパイルして、実行ファイルを作成する
RUN go mod download

# main.go ファイルをコンパイルして、実行ファイルを作成する
RUN go build -o main /app/main.go

# コンテナの8080ポートを公開
EXPOSE 8080

# コンテナ起動後に、実行ファイルを実行する
CMD /app/main
