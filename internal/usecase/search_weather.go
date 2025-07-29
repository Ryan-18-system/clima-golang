package usecase

import (
	"log"

	"github.com/Ryan-18-system/clima-golang/internal/model/brasilapi"
	"github.com/Ryan-18-system/clima-golang/internal/service"
)

type SearchWeather struct {
	BrasilApiService *service.BrasilApiService
	ConversorService *service.ConversorService
}

func NewSearchWeather(brasilApiService *service.BrasilApiService, conversorService *service.ConversorService) *SearchWeather {
	return &SearchWeather{
		BrasilApiService: brasilApiService,
		ConversorService: conversorService,
	}
}
func (sw *SearchWeather) GetWeatherByCep(cep string) (*brasilapi.Address, error) {
	address, err := sw.BrasilApiService.GetCep(cep)
	log.Printf("Address retrieved: %+v\n", address)
	if err != nil {
		return nil, err
	}

	// Assuming we have a method to get weather data by address
	// weatherData, err := sw.WeatherService.GetWeatherByAddress(address)
	// if err != nil {
	// 	return "", err
	// }

	// Convert temperature if needed
	// convertedTemp := sw.ConversorService.ConverterParafahrenheit(weatherData.Temperature)

	return address, nil // Replace with actual weather data
}
