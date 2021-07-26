package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateClientNetworkResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateClientNetworkResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateClientNetworkResponse struct{}"
	}

	return strings.Join([]string{"UpdateClientNetworkResponse", string(data)}, " ")
}
