package brasilapi

type CityResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"nome"`
	State string `json:"estado"`
}

func (c *CityResponse) ToString() string {
	return c.Name + ", " + c.State
}
