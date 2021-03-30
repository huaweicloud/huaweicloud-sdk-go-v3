package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowSecretVersionResponse struct {
	Version        *Version `json:"version,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ShowSecretVersionResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowSecretVersionResponse struct{}"
	}

	return strings.Join([]string{"ShowSecretVersionResponse", string(data)}, " ")
}
