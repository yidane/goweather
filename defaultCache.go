package goweather

import "time"

type DefaultCache struct {
}

type cacheItem struct {
	value   string
	seconds int64
}

var defaultCache = map[string]*cacheItem{}

func (c DefaultCache) Get(k string) string {
	if v, ok := defaultCache[k]; ok {
		if v.seconds > time.Now().Unix() {
			return v.value
		}
		delete(defaultCache, k)
		return ""
	}
	return ""
}

func (c DefaultCache) Set(k, v string, overtime int64) {
	defaultCache[k] = &cacheItem{
		value:   v,
		seconds: time.Now().Unix() + overtime,
	}
}

//TODO:auto remove values which is overtime
