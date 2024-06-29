FROM golang:1.22-alpine as builder

WORKDIR /app

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN swag init -g ./cmd/api/main.go

RUN go mod init backend && go mod tidy

RUN go build -o main ./cmd/api

FROM alpine:latest

COPY --from=builder /app/main /main
COPY --from=builder /app/.env /.env
COPY --from=builder /app/docs /docs

EXPOSE 8888

CMD ["/main"]