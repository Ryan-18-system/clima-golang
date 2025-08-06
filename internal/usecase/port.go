package usecase

import (
	"context"

	"github.com/Ryan-18-system/clima-golang/internal/model/brasilapi"
)

type BrasilApi interface {
	GetCep(cep string) (*brasilapi.Address, error)
	GetCity(city string) (*brasilapi.CityResponse, error)
	GetWeatherByCodeCity(codeCity int) (*brasilapi.WeatherResponse, error)
	GetCepWithContext(ctx context.Context, cep string) (*brasilapi.Address, error)
	GetCityWithContext(ctx context.Context, city string) (*brasilapi.CityResponse, error)
	GetWeatherByCodeCityWithContext(ctx context.Context, codeCity int) (*brasilapi.WeatherResponse, error)
}
type Conversor interface {
	ConverterParafahrenheit(c float64) float64
	ConverterParakelvin(c float64) float64
}
