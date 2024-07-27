FROM --platform=linux/amd64 golang:1.22.3 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o workshop cmd/api/main.go

FROM --platform=linux/amd64 alpine:latest
RUN mkdir /workspace
WORKDIR /workspace/

COPY --from=builder /app/workshop /workspace/workshop

RUN chmod +x /workspace/workshop

EXPOSE 8080

CMD ["/workspace/workshop"]
