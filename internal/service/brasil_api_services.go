package service

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Ryan-18-system/clima-golang/internal/model/brasilapi"
	"github.com/Ryan-18-system/clima-golang/internal/util"
)

type BrasilApiService struct {
	Url string
}

func (brApi *BrasilApiService) GetCep(cep string) (*brasilapi.Address, error) {
	url := fmt.Sprintf("%s/cep/v2/%s", brApi.Url, cep)
	responseByte, err := executeRequest(url)
	if err != nil {
		return nil, err
	}
	address, err := util.ParseJSONResponse[brasilapi.Address](responseByte)
	if err != nil {
		return nil, err
	}
	return address, nil
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
