package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type AddEngressEipV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o AddEngressEipV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AddEngressEipV2Response struct{}"
	}

	return strings.Join([]string{"AddEngressEipV2Response", string(data)}, " ")
}
