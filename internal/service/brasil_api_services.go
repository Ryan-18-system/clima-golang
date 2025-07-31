package service

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Ryan-18-system/clima-golang/internal/model/brasilapi"
	"github.com/Ryan-18-system/clima-golang/internal/util"
)

type BrasilApiService struct {
	Url string
}

func NewBrasilApiService() *BrasilApiService {
	return &BrasilApiService{
		Url: "https://brasilapi.com.br/api",
	}
}

func (brApi *BrasilApiService) GetCep(cep string) (*brasilapi.Address, error) {
	url := fmt.Sprintf("%s/cep/v2/%s", brApi.Url, cep)
	log.Printf("Requesting Brasil API for CEP: %s\n", cep)
	responseByte, err := executeRequest(url)
	if err != nil {
		log.Printf("Error requesting Brasil API for CEP, ERROR: %s\n", err.Error())
		return nil, err
	}
	address, err := util.ParseJSONResponse[brasilapi.Address](responseByte)
	if err != nil {
		log.Printf("Error parsing Brasil API response, ERROR: %s\n", err.Error())
		return nil, err
	}
	return address, nil
}
func (brApi *BrasilApiService) GetCity(city string) (*brasilapi.CityResponse, error) {
	url := fmt.Sprintf("%s/cptec/v1/cidade/%s", brApi.Url, city)
	log.Printf("Requesting Brasil API for City: %s\n", city)
	responseByte, err := executeRequest(url)
	if err != nil {
		return nil, err
	}
	log.Printf("Response from Brasil API: %s\n", string(responseByte))
	cities, err := util.ParseJSONResponse[[]brasilapi.CityResponse](responseByte)
	if err != nil {
		return nil, err
	}
	if len(*cities) == 0 {
		return nil, fmt.Errorf("no cities found in response")
	}
	return &(*cities)[0], nil
}

func (brApi *BrasilApiService) GetWeatherByCodeCity(codeCity int) (*brasilapi.WeatherResponse, error) {
	url := fmt.Sprintf("%s/cptec/v1/clima/previsao/%d", brApi.Url, codeCity)
	log.Printf("Requesting Brasil API for Weather by Code City: %d\n", codeCity)
	responseByte, err := executeRequest(url)
	if err != nil {
		return nil, err
	}
	log.Printf("Response from Brasil API: %s\n", string(responseByte))
	weatherResponse, err := util.ParseJSONResponse[brasilapi.WeatherResponse](responseByte)
	if err != nil {
		return nil, err
	}
	return weatherResponse, nil
}

func executeRequest(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: %s", resp.Status)
	}
	responseByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return responseByte, nil
}
