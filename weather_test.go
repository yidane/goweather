package goweather

import (
	"fmt"
	"testing"
)

func Test_httpGet(t *testing.T) {
	uri := "http://www.weather.com.cn/data/sk/101190408.html"
	s, err := httpGet(uri)

	if err != nil {
		fmt.Println(err)
		t.Error(err)
	}

	fmt.Println(s)
}

func Benchmark_httpGet(b *testing.B) {
	uri := "http://www.weather.com.cn/data/sk/101190408.html"
	for i := 0; i < b.N; i++ {
		httpGet(uri)
	}
}

func TestGetWeather(t *testing.T) {
	var c Cache
	SetCache(&c)
	r, err := GetWeather("101190408")
	if err != nil {
		fmt.Println(err)
		t.Error(err)
	}

	fmt.Println(r)
}

func BenchmarkGetWeather(b *testing.B) {
	var c Cache
	SetCache(&c)
	for i := 0; i < b.N; i++ {
		GetWeather("101190408")
	}
}
