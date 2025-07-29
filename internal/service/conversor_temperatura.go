package service

type ConversorService struct{}

func NewConversorService() *ConversorService {
	return &ConversorService{}
}

func (c *ConversorService) ConverterParafahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}
func (c *ConversorService) ConverterParakelvin(celsius float64) float64 {
	return celsius + 273.15
}
