package dto

import "fmt"

type TemperatureResponse struct {
	TempCelsius    float64 `json:"temp_C"`
	TempFahrenheit float64 `json:"temp_F"`
	TempKelvin     float64 `json:"temp_K"`
}

func (tr *TemperatureResponse) ToJson() string {
	return `{"temp_C":` + fmt.Sprintf("%f", tr.TempCelsius) +
		`, "temp_F":` + fmt.Sprintf("%f", tr.TempFahrenheit) +
		`, "temp_K":` + fmt.Sprintf("%f", tr.TempKelvin) + `}`
}
