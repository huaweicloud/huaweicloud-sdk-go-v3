package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CheckProjectPermissionForAgencyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CheckProjectPermissionForAgencyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CheckProjectPermissionForAgencyResponse struct{}"
	}

	return strings.Join([]string{"CheckProjectPermissionForAgencyResponse", string(data)}, " ")
}
