package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type AssociateAgencyWithProjectPermissionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AssociateAgencyWithProjectPermissionResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AssociateAgencyWithProjectPermissionResponse struct{}"
	}

	return strings.Join([]string{"AssociateAgencyWithProjectPermissionResponse", string(data)}, " ")
}
