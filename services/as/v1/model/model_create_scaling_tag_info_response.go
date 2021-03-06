package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateScalingTagInfoResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateScalingTagInfoResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateScalingTagInfoResponse struct{}"
	}

	return strings.Join([]string{"CreateScalingTagInfoResponse", string(data)}, " ")
}
