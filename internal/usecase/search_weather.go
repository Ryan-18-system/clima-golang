package usecase

import (
	"errors"
	"log"

	"github.com/Ryan-18-system/clima-golang/internal/model/brasilapi"
	"github.com/Ryan-18-system/clima-golang/internal/model/dto"
)

type SearchWeather struct {
	BrasilApiService BrasilApi
	ConversorService Conversor
}

func NewSearchWeather(brasilApiService BrasilApi, conversorService Conversor) *SearchWeather {
	return &SearchWeather{
		BrasilApiService: brasilApiService,
		ConversorService: conversorService,
	}
}

func (sw *SearchWeather) GetWeatherByCep(cep string) (*dto.TemperatureResponse, error) {
	address, err := sw.BrasilApiService.GetCep(cep)
	log.Printf("Address retrieved: %+v\n", address)
	if err != nil {
		return nil, err
	}
	log.Printf("City found: %s, State: %s\n", address.City, address.State)
	cityResponse, err := sw.BrasilApiService.GetCity(address.City)
	if err != nil {
		log.Printf("City not found for CEP: %s\n", cep)
		return nil, err
	}
	weatherResponse, err := sw.BrasilApiService.GetWeatherByCodeCity(cityResponse.ID)
	if err != nil {
		log.Printf("Weather not found for city: %s, ID: %d\n", cityResponse.Name, cityResponse.ID)
		return nil, err
	}
	tempResponse, err := sw.MapTemperatures(weatherResponse)
	if err != nil {
		log.Printf("Error mapping temperatures for city: %s, ID: %d\n", cityResponse.Name, cityResponse.ID)
		return nil, err
	}
	return tempResponse, nil
}

// Mapper para calcular Fahrenheit e Kelvin
func (sw *SearchWeather) MapTemperatures(weather *brasilapi.WeatherResponse) (*dto.TemperatureResponse, error) {
	var tempResponse dto.TemperatureResponse
	if weather == nil {
		return nil, errors.New("weather data is nil")
	}
	if len(weather.Clima) == 0 {
		return nil, errors.New("weather climate data is nil")
	}
	celsius := float64(weather.Clima[0].Max)
	tempResponse.TempCelsius = float64(celsius)
	tempResponse.TempFahrenheit = sw.ConversorService.ConverterParafahrenheit(celsius)
	tempResponse.TempKelvin = sw.ConversorService.ConverterParakelvin(celsius)
	return &tempResponse, nil
}
