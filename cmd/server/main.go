package main

import (
	"net/http"

	"github.com/Ryan-18-system/clima-golang/internal/service"
)

func main() {
	url := "https://brasilapi.com.br/api"
	brasilApiService := &service.BrasilApiService{Url: url}
	cep := "01001-000"
	address, err := brasilApiService.GetCep(cep)
	if err != nil {
		panic(err)
	}
	println("Address:", address.FormatedAddressBrasilApi())
	mux := http.NewServeMux()
	mux.HandleFunc("POST /temperatura/cep", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Material created successfully"))
	})
	http.ListenAndServe(":8080", mux)
}
