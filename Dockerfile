FROM golang:alpine

WORKDIR /app

COPY . .

RUN go get ./...

RUN go build -v -o transfer-app /app/cmd/main.go

EXPOSE 9090

CMD ["/app/transfer-app"]