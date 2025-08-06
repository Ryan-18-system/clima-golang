package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/Ryan-18-system/clima-golang/internal/di"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

func initTracer() func() {
	zipkinURL := os.Getenv("ZIPKIN_URL")
	if zipkinURL == "" {
		zipkinURL = "http://localhost:9411/api/v2/spans"
	}
	exporter, err := zipkin.New(
		zipkinURL,
	)
	if err != nil {
		log.Fatalf("failed to create zipkin exporter: %v", err)
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("clima-golang-api"),
		)),
	)
	otel.SetTracerProvider(tp)
	return func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}
}
func main() {
	shutdown := initTracer()
	defer shutdown()

	climaController := di.InitializeClimaController()
	mux := http.NewServeMux()
	mux.HandleFunc("POST /temperatura/cep", climaController.SearchWeatherByZipCode)
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Boas-vindas a api de clima feita em golang"))
	})
	log.Println("Servidor iniciado na porta 8080")
	http.ListenAndServe(":8080", mux)
}
