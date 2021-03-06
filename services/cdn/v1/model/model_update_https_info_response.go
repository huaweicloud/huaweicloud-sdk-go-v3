package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateHttpsInfoResponse struct {
	Https          *HttpInfoResponseBody `json:"https,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o UpdateHttpsInfoResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateHttpsInfoResponse struct{}"
	}

	return strings.Join([]string{"UpdateHttpsInfoResponse", string(data)}, " ")
}
