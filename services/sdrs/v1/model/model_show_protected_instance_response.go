package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowProtectedInstanceResponse struct {
	ProtectedInstance *ShowProtectedInstanceParams `json:"protected_instance,omitempty"`
	HttpStatusCode    int                          `json:"-"`
}

func (o ShowProtectedInstanceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowProtectedInstanceResponse struct{}"
	}

	return strings.Join([]string{"ShowProtectedInstanceResponse", string(data)}, " ")
}
