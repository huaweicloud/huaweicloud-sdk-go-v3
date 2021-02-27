package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateConfigurationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateConfigurationResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateConfigurationResponse struct{}"
	}

	return strings.Join([]string{"UpdateConfigurationResponse", string(data)}, " ")
}
