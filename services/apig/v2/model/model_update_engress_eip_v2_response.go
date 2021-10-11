package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateEngressEipV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateEngressEipV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateEngressEipV2Response struct{}"
	}

	return strings.Join([]string{"UpdateEngressEipV2Response", string(data)}, " ")
}
