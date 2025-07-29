package di

import (
	"github.com/Ryan-18-system/clima-golang/internal/adapter/http/controller"
	"github.com/Ryan-18-system/clima-golang/internal/service"
	"github.com/Ryan-18-system/clima-golang/internal/usecase"
)

func InitializeClimaController() *controller.ClimaController {
	brasilApiService := service.NewBrasilApiService()
	conversorService := service.NewConversorService()
	searchWeather := usecase.NewSearchWeather(brasilApiService, conversorService)
	return controller.NewClimaController(searchWeather)
}
