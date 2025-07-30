package usecase

import "github.com/Ryan-18-system/clima-golang/internal/model/brasilapi"

type BrasilApi interface {
	GetCep(cep string) (*brasilapi.Address, error)
	GetCity(city string) (*brasilapi.CityResponse, error)
	GetWeatherByCodeCity(codeCity int) (*brasilapi.WeatherResponse, error)
}
type Conversor interface {
	ConverterParafahrenheit(c float64) float64
	ConverterParakelvin(c float64) float64
}
