package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteEnvironmentResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteEnvironmentResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteEnvironmentResponse struct{}"
	}

	return strings.Join([]string{"DeleteEnvironmentResponse", string(data)}, " ")
}
