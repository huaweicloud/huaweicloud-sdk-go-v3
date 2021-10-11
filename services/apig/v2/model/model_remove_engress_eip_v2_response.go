package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RemoveEngressEipV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o RemoveEngressEipV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RemoveEngressEipV2Response struct{}"
	}

	return strings.Join([]string{"RemoveEngressEipV2Response", string(data)}, " ")
}
