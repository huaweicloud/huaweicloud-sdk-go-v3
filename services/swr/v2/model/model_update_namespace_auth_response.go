package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateNamespaceAuthResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateNamespaceAuthResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateNamespaceAuthResponse struct{}"
	}

	return strings.Join([]string{"UpdateNamespaceAuthResponse", string(data)}, " ")
}
