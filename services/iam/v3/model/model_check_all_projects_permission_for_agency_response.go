package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CheckAllProjectsPermissionForAgencyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CheckAllProjectsPermissionForAgencyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CheckAllProjectsPermissionForAgencyResponse struct{}"
	}

	return strings.Join([]string{"CheckAllProjectsPermissionForAgencyResponse", string(data)}, " ")
}
