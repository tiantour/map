package qq

import (
	"encoding/json"
	"fmt"

	"github.com/tiantour/fetch"
)

// Key map api key
var Key = ""

// QQ qq
type QQ struct{}

// NewQQ new qq
func NewQQ() *QQ {
	return &QQ{}
}

type (
	// Response response
	Response struct {
		Status  int
		Message string
		Result  Result
	}

	// Result result
	Result struct {
		Title             string
		Similarity        float64
		Deviation         int
		Reliability       int
		Address           string
		Location          Location
		AddressComponents Address
	}

	// Address address
	Address struct {
		Province     string
		City         string
		District     string
		Street       string
		StreetNumber string
	}

	// Location location
	Location struct {
		Lng float64
		Lat float64
	}
)

// ToAddress location to address
func (q QQ) ToAddress(lng, lat float64) (Response, error) {
	result := Response{}
	body, err := fetch.Cmd(fetch.Request{
		Method: "GET",
		URL: fmt.Sprintf("http://apis.map.qq.com/ws/geocoder/v1/?location=%f,%f&get_poi=1&output=json&callback=cb&key=%s",
			lat,
			lng,
			Key,
		),
	})
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	return result, err
}

// ToLocation address to location
func (q QQ) ToLocation(address string) (Response, error) {
	result := Response{}
	body, err := fetch.Cmd(fetch.Request{
		Method: "GET",
		URL: fmt.Sprintf("http://apis.map.qq.com/ws/geocoder/v1/?address=%s&output=json&callback=cb&key=%s",
			address,
			Key,
		),
	})
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	return result, err
}
