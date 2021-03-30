package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteSecretResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSecretResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteSecretResponse struct{}"
	}

	return strings.Join([]string{"DeleteSecretResponse", string(data)}, " ")
}
