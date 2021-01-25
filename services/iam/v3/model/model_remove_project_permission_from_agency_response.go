package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RemoveProjectPermissionFromAgencyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RemoveProjectPermissionFromAgencyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RemoveProjectPermissionFromAgencyResponse struct{}"
	}

	return strings.Join([]string{"RemoveProjectPermissionFromAgencyResponse", string(data)}, " ")
}
