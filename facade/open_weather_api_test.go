package facade

import (
	"../facade"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOpenWeatherMap_responseParser(t *testing.T) {
	r := facade.GetMockData()
	openWeatherMap := facade.CurrentWeatherData{APIkey: "1680a7cc01fef9a05aab87514f2216dc"}
	weather, err := openWeatherMap.ResponseParser(r)
	if err != nil {
		t.Fatal(err)
	}
	assert.EqualValues(t, 0, weather.ID)
}
func TestCurrentWeatherData_GetByCityAndCountryCode(t *testing.T) {
	weatherMap := facade.CurrentWeatherData{APIkey: "1680a7cc01fef9a05aab87514f2216dc"}
	weather, err := weatherMap.GetByCityAndCountryCode("London", "uk")
	if err != nil {
		t.Fatal(err)
	}
	assert.EqualValues(t, 2643743, weather.ID)
}
