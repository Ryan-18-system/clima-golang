FROM golang:latest AS builder

WORKDIR /app
COPY . .

# Corrigido: aponta para o diretório onde está o main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o climaapi ./cmd/server

FROM scratch
WORKDIR /app
COPY --from=builder /app/climaapi .

ENTRYPOINT ["./climaapi"]
