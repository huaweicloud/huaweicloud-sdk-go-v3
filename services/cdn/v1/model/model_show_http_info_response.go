package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowHttpInfoResponse struct {
	Https          *HttpInfoResponseBody `json:"https,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ShowHttpInfoResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowHttpInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowHttpInfoResponse", string(data)}, " ")
}
