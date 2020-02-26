package facade

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type CurrentWeatherDataRetriever interface {
	GetByCityAndCountryCode(city, countryCode string) (*Weather, error)
	GetByGeoCoord(lat, lon float32) (*Weather, error)
}

type Weather struct {
	Base   string `json:"base"`
	Clouds struct {
		All float64 `json:"all"`
	} `json:"clouds"`
	Cod   float64 `json:"cod"`
	Coord struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"coord"`
	Dt   float64 `json:"dt"`
	ID   float64 `json:"id"`
	Main struct {
		Humidity float64 `json:"humidity"`
		Pressure float64 `json:"pressure"`
		Temp     float64 `json:"temp"`
		TempMax  float64 `json:"temp_max"`
		TempMin  float64 `json:"temp_min"`
	} `json:"main"`
	Name string `json:"name"`
	Sys  struct {
		Country string  `json:"country"`
		ID      float64 `json:"id"`
		Message float64 `json:"message"`
		Sunrise float64 `json:"sunrise"`
		Sunset  float64 `json:"sunset"`
		Type    float64 `json:"type"`
	} `json:"sys"`
	Visibility float64 `json:"visibility"`
	Weather    []struct {
		Description string  `json:"description"`
		Icon        string  `json:"icon"`
		ID          float64 `json:"id"`
		Main        string  `json:"main"`
	} `json:"weather"`
	Wind struct {
		Deg   float64 `json:"deg"`
		Gust  float64 `json:"gust"`
		Speed float64 `json:"speed"`
	} `json:"wind"`
}

type CurrentWeatherData struct {
	APIkey string
}

func (p *CurrentWeatherData) ResponseParser(body io.Reader) (*Weather, error) {
	var response Weather
	err := json.NewDecoder(body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func GetMockData() io.Reader {
	response := `{"coord":{"lon":-122.09,"lat":37.39},"weather":[{"id":500,"main":"Rain","description":"light rain","icon":"10d"}],"base":"stations","main":{"temp":280.44,"pressure":1017,"humidity":61,"temp_min":279.15,"temp_max":281.15},"visibility":12874,"wind":{"speed":8.2,"deg":340,"gust":11.3},"clouds":{"all":1},"dt":1519061700,"sys":{"type":1,"id":392,"message":0.0027,"country":"US","sunrise":1519051894,"sunset":1519091585},"id":0,"name":"Mountain View","cod":200}`
	r := bytes.NewReader([]byte(response))
	return r

}

func (c *CurrentWeatherData) GetByGeoCoord(lat, lon float32) (*Weather, error) {
	return c.doRequest(fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&APPID=%s", lat, lon, c.APIkey))
}

func (c *CurrentWeatherData) GetByCityAndCountryCode(city, countryCode string) (*Weather, error) {

	return c.doRequest(fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s,%s&APPID=%s", city, countryCode, c.APIkey))
}
func (o *CurrentWeatherData) doRequest(uri string) (*Weather, error) {
	var weather *Weather
	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, erro := client.Do(req)
	if erro != nil {
		return nil, erro
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		byt, errMsg := ioutil.ReadAll(resp.Body)
		if errMsg == nil {
			errMsg = fmt.Errorf("%s", string(byt))
		}
		err = fmt.Errorf("Status code was %d, aborting. Error message was:\n%s\n", resp.StatusCode, errMsg)
		return nil, err
	}
	weather, err = o.ResponseParser(resp.Body)
	return weather, nil
}
