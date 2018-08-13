package goweather

import (
	"fmt"
	"testing"
)

func Test_loadCities(t *testing.T) {
	LoadCities()

	v, err := CityCache.Get("101010100")
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(v)

	fmt.Println(GetWeather("101010100"))
}
