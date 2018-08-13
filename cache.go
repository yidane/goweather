package goweather

import (
	"github.com/Tencent/bk-cmdb/src/framework/core/errors"
	"io/ioutil"
	"net/http"
)

var WeatherCache Cache
var CityCache Cache
var defaultOverTime int64 = 1800 // seconds of half-hour

var errKeyNotFound = errors.New("the key matches nothing")
var errKeyExpired = errors.New("the key is expired")

type Cache interface {
	Get(k string) (string, error)
	Set(k, v string, timeSpan int64) error
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

func RegisterCache(weatherCache, cityCache *Cache) {
	if *weatherCache == nil {
		panic("WeatherCache cannot be null")
	}

	if *cityCache == nil {
		panic("CityCache cannot be null")
	}

	WeatherCache = *weatherCache
	CityCache = *cityCache
}

func init() {
	var wc Cache = DefaultWeatherCache{}
	var cc Cache = DefaultCityCache{}
	RegisterCache(&wc, &cc)
}
