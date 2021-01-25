package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RemoveAllProjectsPermissionFromAgencyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RemoveAllProjectsPermissionFromAgencyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RemoveAllProjectsPermissionFromAgencyResponse struct{}"
	}

	return strings.Join([]string{"RemoveAllProjectsPermissionFromAgencyResponse", string(data)}, " ")
}
