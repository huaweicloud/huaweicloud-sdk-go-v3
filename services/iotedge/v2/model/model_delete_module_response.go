package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteModuleResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteModuleResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteModuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteModuleResponse", string(data)}, " ")
}
