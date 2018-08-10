package goweather

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const getWeatherURI = "http://www.weather.com.cn/data/sk/%s.html"

func GetCity() {

}

func GetWeather(code string) (string, error) {
	cr := CurrentCache.Get(code)
	if len(cr) > 0 {
		return cr, nil
	}

	rr, err := httpGet(fmt.Sprintf(getWeatherURI, code))
	if err != nil {
		return "", err
	}

	CurrentCache.Set(code, rr, DefaultOverTime)
	return rr, nil
}

func httpGet(uri string) (string, error) {
	req, _ := http.NewRequest("GET", uri, nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	bs, err := ioutil.ReadAll(res.Body)

	return string(bs), nil
}
