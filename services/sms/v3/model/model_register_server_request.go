package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RegisterServerRequest struct {
	Body *PostSourceServerBody `json:"body,omitempty"`
}

func (o RegisterServerRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RegisterServerRequest struct{}"
	}

	return strings.Join([]string{"RegisterServerRequest", string(data)}, " ")
}
