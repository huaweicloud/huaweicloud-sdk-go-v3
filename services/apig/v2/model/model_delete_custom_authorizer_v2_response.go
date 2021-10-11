package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteCustomAuthorizerV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteCustomAuthorizerV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteCustomAuthorizerV2Response struct{}"
	}

	return strings.Join([]string{"DeleteCustomAuthorizerV2Response", string(data)}, " ")
}
