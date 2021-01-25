package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type AssociateAgencyWithAllProjectsPermissionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AssociateAgencyWithAllProjectsPermissionResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AssociateAgencyWithAllProjectsPermissionResponse struct{}"
	}

	return strings.Join([]string{"AssociateAgencyWithAllProjectsPermissionResponse", string(data)}, " ")
}
