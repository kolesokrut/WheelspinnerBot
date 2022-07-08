package banking

type NationalBank struct {
	R030         int     `json:"r030"`
	Txt          string  `json:"txt"`
	Rate         float64 `json:"rate"`
	Cc           string  `json:"cc"`
	Exchangedate string  `json:"exchangedate"`
}

type OpenWeather struct {
	Message string `json:"message"`
	Cod     string `json:"cod"`
	Count   int    `json:"count"`
	List    []struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Coord struct {
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		} `json:"coord"`
		Main struct {
			Temp      int     `json:"temp"`
			FeelsLike float64 `json:"feels_like"`
			TempMin   int     `json:"temp_min"`
			TempMax   int     `json:"temp_max"`
			Pressure  int     `json:"pressure"`
			Humidity  int     `json:"humidity"`
			SeaLevel  int     `json:"sea_level"`
			GrndLevel int     `json:"grnd_level"`
		} `json:"main"`
		Dt   int `json:"dt"`
		Wind struct {
			Speed float64 `json:"speed"`
			Deg   int     `json:"deg"`
		} `json:"wind"`
		Sys struct {
			Country string `json:"country"`
		} `json:"sys"`
		Rain   interface{} `json:"rain"`
		Snow   interface{} `json:"snow"`
		Clouds struct {
			All int `json:"all"`
		} `json:"clouds"`
		Weather []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
	} `json:"list"`
}

type PrivatBank24 struct {
	Ccy     string  `json:"ccy"`
	BaseCcy string  `json:"base_ccy"`
	Buy     float64 `json:",string"`
	Sale    float64 `json:",string"`
}
