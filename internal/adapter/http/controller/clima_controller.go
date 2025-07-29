package controller

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/Ryan-18-system/clima-golang/internal/model/dto"
	"github.com/Ryan-18-system/clima-golang/internal/usecase"
)

type ClimaController struct {
	SearchWeatherService *usecase.SearchWeather
}

func NewClimaController(searchWeatherService *usecase.SearchWeather) *ClimaController {
	return &ClimaController{
		SearchWeatherService: searchWeatherService,
	}
}
func (c *ClimaController) SearchWeatherByZipCode(w http.ResponseWriter, r *http.Request) {
	var req dto.CepRequest
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validação de formato de CEP: 5 dígitos + hífen + 3 dígitos OU 8 dígitos
	validFormat := regexp.MustCompile(`^\d{5}-?\d{3}$`).MatchString
	if !validFormat(req.Cep) {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	result, err := c.SearchWeatherService.GetWeatherByCep(req.Cep)
	if err != nil {
		http.Error(w, "can not find zipcode", 404)
		return

		// http.Error(w, err.Error(), http.StatusInternalServerError)
		// return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"result":"` + result.Estado + `"}`))
}
