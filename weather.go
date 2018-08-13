package goweather

import (
	"fmt"
)

const getWeatherURI = "http://www.weather.com.cn/data/sk/%s.html"

func GetWeather(code string) (string, error) {
	cr, err := WeatherCache.Get(code)
	if err != nil && err != errKeyNotFound && err != errKeyExpired {
		return cr, err
	}

	return getWeatherByURI(code)
}

func getWeatherByURI(code string) (string, error) {
	rr, err := httpGet(fmt.Sprintf(getWeatherURI, code))
	if err != nil {
		return "", err
	}

	WeatherCache.Set(code, rr, defaultOverTime)
	return rr, nil
}
