package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneDeleteProtocolResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o KeystoneDeleteProtocolResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneDeleteProtocolResponse struct{}"
	}

	return strings.Join([]string{"KeystoneDeleteProtocolResponse", string(data)}, " ")
}
