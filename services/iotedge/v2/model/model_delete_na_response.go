package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteNaResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteNaResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteNaResponse struct{}"
	}

	return strings.Join([]string{"DeleteNaResponse", string(data)}, " ")
}
