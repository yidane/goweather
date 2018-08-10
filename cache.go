package goweather

var CurrentCache Cache
var DefaultOverTime int64 = 1800 // seconds of half-hour

type Cache interface {
	Get(k string) string
	Set(k, v string, overTime int64)
}

func SetCache(c *Cache) {
	if *c == nil {
		*c = DefaultCache{}
	}

	CurrentCache = *c
}
