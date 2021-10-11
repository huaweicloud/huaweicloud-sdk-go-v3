package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteExternalEntityResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteExternalEntityResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteExternalEntityResponse struct{}"
	}

	return strings.Join([]string{"DeleteExternalEntityResponse", string(data)}, " ")
}
