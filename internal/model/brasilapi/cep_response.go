package brasilapi

type Address struct {
	Cep          string `json:"cep"`
	Logradouro   string `json:"logradouro"`
	Complemento  string `json:"complemento"`
	Unidade      string `json:"unidade"`
	Bairro       string `json:"bairro"`
	Localidade   string `json:"localidade"`
	Uf           string `json:"uf"`
	Estado       string `json:"estado"`
	Regiao       string `json:"regiao"`
	Ibge         string `json:"ibge"`
	Gia          string `json:"gia"`
	Ddd          string `json:"ddd"`
	Siafi        string `json:"siafi"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
	Location     struct {
		Type        string `json:"type"`
		Coordinates struct {
			Longitude string `json:"longitude"`
			Latitude  string `json:"latitude"`
		} `json:"coordinates"`
	} `json:"location"`
}

func (brasilApi *Address) FormatedAddressBrasilApi() string {
	return brasilApi.Street + ", " + brasilApi.Neighborhood + ", " + brasilApi.City + " - " + brasilApi.State + ", " + brasilApi.Cep
}
