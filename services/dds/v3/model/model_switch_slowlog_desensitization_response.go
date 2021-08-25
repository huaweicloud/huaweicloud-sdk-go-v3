package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type SwitchSlowlogDesensitizationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SwitchSlowlogDesensitizationResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SwitchSlowlogDesensitizationResponse struct{}"
	}

	return strings.Join([]string{"SwitchSlowlogDesensitizationResponse", string(data)}, " ")
}
