package main

import (
	"net/http"

	"github.com/Ryan-18-system/clima-golang/internal/di"
)

func main() {
	climaController := di.InitializeClimaController()
	mux := http.NewServeMux()
	mux.HandleFunc("POST /temperatura/cep", climaController.SearchWeatherByZipCode)
	http.ListenAndServe(":8080", mux)
}
