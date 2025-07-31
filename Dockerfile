FROM golang:latest AS builder

WORKDIR /app

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o climaapi


FROM scratch
WORKDIR /app
COPY --from=build /app/climaapi .

ENTRYPOINT ["./climaapi"]
