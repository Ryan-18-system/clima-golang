package model

type WeatherResponse struct {
	Data     WeatherData `json:"data"`
	Location Location    `json:"location"`
}

type WeatherData struct {
	Time   string        `json:"time"`
	Values WeatherValues `json:"values"`
}

type WeatherValues struct {
	CloudBase                float64 `json:"cloudBase"`
	CloudCeiling             float64 `json:"cloudCeiling"`
	CloudCover               int     `json:"cloudCover"`
	DewPoint                 float64 `json:"dewPoint"`
	FreezingRainIntensity    float64 `json:"freezingRainIntensity"`
	Humidity                 int     `json:"humidity"`
	PrecipitationProbability int     `json:"precipitationProbability"`
	PressureSurfaceLevel     float64 `json:"pressureSurfaceLevel"`
	RainIntensity            float64 `json:"rainIntensity"`
	SleetIntensity           float64 `json:"sleetIntensity"`
	SnowIntensity            float64 `json:"snowIntensity"`
	Temperature              float64 `json:"temperature"`
	TemperatureApparent      float64 `json:"temperatureApparent"`
	UvHealthConcern          int     `json:"uvHealthConcern"`
	UvIndex                  int     `json:"uvIndex"`
	Visibility               float64 `json:"visibility"`
	WeatherCode              int     `json:"weatherCode"`
	WindDirection            int     `json:"windDirection"`
	WindGust                 float64 `json:"windGust"`
	WindSpeed                float64 `json:"windSpeed"`
}

type Location struct {
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
	Name string  `json:"name"`
	Type string  `json:"type"`
}
