package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteEnvironmentV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteEnvironmentV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteEnvironmentV2Response struct{}"
	}

	return strings.Join([]string{"DeleteEnvironmentV2Response", string(data)}, " ")
}
