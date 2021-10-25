package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateProtectedInstanceNameResponse struct {
	ProtectedInstance *ShowProtectedInstanceParams `json:"protected_instance,omitempty"`
	HttpStatusCode    int                          `json:"-"`
}

func (o UpdateProtectedInstanceNameResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateProtectedInstanceNameResponse struct{}"
	}

	return strings.Join([]string{"UpdateProtectedInstanceNameResponse", string(data)}, " ")
}
