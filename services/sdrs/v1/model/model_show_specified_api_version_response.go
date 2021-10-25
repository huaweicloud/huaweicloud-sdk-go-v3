package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowSpecifiedApiVersionResponse struct {
	Version        *ShowApiVersionParams `json:"version,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ShowSpecifiedApiVersionResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowSpecifiedApiVersionResponse struct{}"
	}

	return strings.Join([]string{"ShowSpecifiedApiVersionResponse", string(data)}, " ")
}
