package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RegisterServerAutoRecoveryResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RegisterServerAutoRecoveryResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RegisterServerAutoRecoveryResponse struct{}"
	}

	return strings.Join([]string{"RegisterServerAutoRecoveryResponse", string(data)}, " ")
}
