package brasilapi

type WeatherResponse struct {
	Cidade       string        `json:"cidade"`
	Estado       string        `json:"estado"`
	AtualizadoEm string        `json:"atualizado_em"`
	Clima        []PrevisaoDia `json:"clima"`
}

type PrevisaoDia struct {
	Data         string `json:"data"`
	Condicao     string `json:"condicao"`
	CondicaoDesc string `json:"condicao_desc"`
	Min          int    `json:"min"`
	Max          int    `json:"max"`
	IndiceUV     int    `json:"indice_uv"`
}
