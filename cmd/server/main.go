package main

import (
	"net/http"

	"github.com/Ryan-18-system/clima-golang/internal/di"
)

func main() {
	climaController := di.InitializeClimaController()
	mux := http.NewServeMux()
	mux.HandleFunc("POST /temperatura/cep", climaController.SearchWeatherByZipCode)
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Boas-vindas a api de clima feita em golang"))
	})
	http.ListenAndServe(":8080", mux)
}
