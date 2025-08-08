package usecase

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/Ryan-18-system/clima-golang/internal/model/brasilapi"
	"github.com/Ryan-18-system/clima-golang/internal/model/dto"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
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

// Novo método com contexto para tracing distribuído
func (sw *SearchWeather) GetWeatherByCepWithContext(ctx context.Context, cep string) (*dto.TemperatureResponse, error) {
	tracer := otel.Tracer("search-weather")
	ctx, span := tracer.Start(ctx, "Serviço B - Orquestração")
	defer span.End()

	// Span para busca de CEP
	cepCtx, cepSpan := tracer.Start(ctx, "Busca CEP")
	startCep := time.Now()
	address, err := sw.BrasilApiService.GetCepWithContext(cepCtx, cep)

	if err != nil {
		return nil, err
	}
	cepSpan.SetAttributes(attribute.String("cep", cep))
	cepSpan.SetAttributes(attribute.String("city", address.City))
	cepSpan.SetAttributes(attribute.String("state", address.State))
	cepSpan.SetAttributes(attribute.Float64("duration_ms", float64(time.Since(startCep).Milliseconds())))
	cepSpan.End()
	log.Printf("Address retrieved: %+v\n", address)
	// Span para busca de cidade
	cityCtx, citySpan := tracer.Start(ctx, "Busca Cidade")
	cityResponse, err := sw.BrasilApiService.GetCityWithContext(cityCtx, address.City)
	if err != nil {
		log.Printf("City not found for CEP: %s\n", cep)
		return nil, err
	}
	citySpan.SetAttributes(attribute.String("city", address.City))
	citySpan.End()

	// Span para busca de clima
	weatherCtx, weatherSpan := tracer.Start(ctx, "Busca Clima")
	weatherResponse, err := sw.BrasilApiService.GetWeatherByCodeCityWithContext(weatherCtx, cityResponse.ID)
	if err != nil {
		log.Printf("Weather not found for city: %s, ID: %d\n", cityResponse.Name, cityResponse.ID)
		return nil, err
	}
	weatherSpan.SetAttributes(attribute.Int("city_id", cityResponse.ID))
	weatherSpan.End()

	tempResponse, err := sw.MapTemperatures(weatherResponse)
	if err != nil {
		log.Printf("Error mapping temperatures for city: %s, ID: %d\n", cityResponse.Name, cityResponse.ID)
		return nil, err
	}
	return tempResponse, nil
}

// Métodos originais continuam para compatibilidade
func (sw *SearchWeather) GetWeatherByCep(cep string) (*dto.TemperatureResponse, error) {
	return sw.GetWeatherByCepWithContext(context.Background(), cep)
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
