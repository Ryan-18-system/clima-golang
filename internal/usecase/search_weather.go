package usecase

import "github.com/Ryan-18-system/clima-golang/internal/service"

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
func (sw *SearchWeather) GetWeatherByCep(cep string) (string, error) {
	address, err := sw.BrasilApiService.GetCep(cep)
	if err != nil {
		return "", err
	}

	// Assuming we have a method to get weather data by address
	// weatherData, err := sw.WeatherService.GetWeatherByAddress(address)
	// if err != nil {
	// 	return "", err
	// }

	// Convert temperature if needed
	// convertedTemp := sw.ConversorService.ConverterParafahrenheit(weatherData.Temperature)

	return address.FormatedAddressBrasilApi(), nil // Replace with actual weather data
}
