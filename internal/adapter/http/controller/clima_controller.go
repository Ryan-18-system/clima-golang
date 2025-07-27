package controller

import (
	"net/http"

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

}
