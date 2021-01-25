package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteApplicationConfigurationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteApplicationConfigurationResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteApplicationConfigurationResponse struct{}"
	}

	return strings.Join([]string{"DeleteApplicationConfigurationResponse", string(data)}, " ")
}
