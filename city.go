package goweather

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"
)

const getCityURI = "http://mobile.weather.com.cn/js/citylist.xml"

type Location struct {
	ID         string
	Name       string
	EnName     string
	ParentName string
}

func (l Location) String() string {
	b, _ := json.Marshal(l)
	return string(b)
}

func GetCity(code string) string {
	v, err := CityCache.Get(code)
	if err != nil {
		return err.Error()
	}

	return v
}

func LoadCities() error {
	r, err := httpGet(getCityURI)
	if err != nil {
		fmt.Println(err)
		return err
	}

	xmlDecoder := xml.NewDecoder(strings.NewReader(r))
	for t, err := xmlDecoder.Token(); err == nil; t, err = xmlDecoder.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			name := token.Name.Local
			if name == "d" {
				l := Location{}
				for _, attr := range token.Attr {
					switch attr.Name.Local {
					case "d1":
						l.ID = attr.Value
					case "d2":
						l.Name = attr.Value
					case "d3":
						l.EnName = attr.Value
					case "d4":
						l.ParentName = attr.Value
					}
				}

				CityCache.Set(l.ID, l.String(), defaultOverTime)
			}
		}
	}

	return nil
}
