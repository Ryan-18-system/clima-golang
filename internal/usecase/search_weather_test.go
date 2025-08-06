package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/Ryan-18-system/clima-golang/internal/model/brasilapi"
	"github.com/Ryan-18-system/clima-golang/internal/usecase"
	"github.com/stretchr/testify/assert"
)

// --- Mocks das interfaces ---

type mockBrasilApi struct{}

func (m *mockBrasilApi) GetCep(cep string) (*brasilapi.Address, error) {
	if cep == "00000000" {
		return nil, errors.New("invalid cep")
	}
	return &brasilapi.Address{City: "João Pessoa", State: "PB"}, nil
}

func (m *mockBrasilApi) GetCity(city string) (*brasilapi.CityResponse, error) {
	if city == "Desconhecida" {
		return nil, errors.New("city not found")
	}
	return &brasilapi.CityResponse{ID: 101, Name: city}, nil
}

func (m *mockBrasilApi) GetWeatherByCodeCity(cityID int) (*brasilapi.WeatherResponse, error) {
	return &brasilapi.WeatherResponse{
		Cidade: "João Pessoa",
		Clima: []brasilapi.PrevisaoDia{
			{Max: 30},
		},
	}, nil
}
func (m *mockBrasilApi) GetCepWithContext(ctx context.Context, cep string) (*brasilapi.Address, error) {
	return m.GetCep(cep)
}
func (m *mockBrasilApi) GetCityWithContext(ctx context.Context, city string) (*brasilapi.CityResponse, error) {
	return m.GetCity(city)
}
func (m *mockBrasilApi) GetWeatherByCodeCityWithContext(ctx context.Context, cityID int) (*brasilapi.WeatherResponse, error) {
	return m.GetWeatherByCodeCity(cityID)
}

type mockConversor struct{}

func (m *mockConversor) ConverterParafahrenheit(c float64) float64 {
	return (c * 1.8) + 32
}
func (m *mockConversor) ConverterParakelvin(c float64) float64 {
	return c + 273.15
}

func TestGetWeatherByCep_Success(t *testing.T) {
	sw := usecase.NewSearchWeather(&mockBrasilApi{}, &mockConversor{})

	result, err := sw.GetWeatherByCep("58000000")
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 30.0, result.TempCelsius)
	assert.Equal(t, 86.0, result.TempFahrenheit)
	assert.Equal(t, 303.15, result.TempKelvin)
}

func TestGetWeatherByCep_InvalidCep(t *testing.T) {
	sw := usecase.NewSearchWeather(&mockBrasilApi{}, &mockConversor{})

	result, err := sw.GetWeatherByCep("00000000")
	assert.Error(t, err)
	assert.Nil(t, result)
}
