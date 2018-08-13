package goweather

import (
	"time"
)

type DefaultWeatherCache struct {
}

type DefaultCityCache struct {
}

type cacheItem struct {
	value   string
	seconds int64
}

var defaultWeatherCache = map[string]*cacheItem{}
var defaultCityCache = map[string]*cacheItem{}

func (c DefaultWeatherCache) Get(k string) (string, error) {
	if v, ok := defaultWeatherCache[k]; ok {
		if v.seconds > time.Now().Unix() {
			return v.value, nil
		}
		delete(defaultWeatherCache, k)
		return ``, errKeyExpired
	}
	return "", errKeyNotFound
}

func (c DefaultWeatherCache) Set(k, v string, overtime int64) error {
	defaultWeatherCache[k] = &cacheItem{
		value:   v,
		seconds: time.Now().Unix() + overtime,
	}

	return nil
}

func (c DefaultCityCache) Get(k string) (string, error) {
	if v, ok := defaultCityCache[k]; ok {
		if v.seconds > time.Now().Unix() {
			return v.value, nil
		}
		delete(defaultCityCache, k)
		return "", errKeyExpired
	}
	return "", errKeyNotFound
}

func (c DefaultCityCache) Set(k, v string, overtime int64) error {
	defaultCityCache[k] = &cacheItem{
		value:   v,
		seconds: time.Now().Unix() + overtime,
	}
	return nil
}
