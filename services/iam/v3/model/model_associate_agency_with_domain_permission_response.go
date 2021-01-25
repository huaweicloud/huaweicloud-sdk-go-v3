package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type AssociateAgencyWithDomainPermissionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AssociateAgencyWithDomainPermissionResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AssociateAgencyWithDomainPermissionResponse struct{}"
	}

	return strings.Join([]string{"AssociateAgencyWithDomainPermissionResponse", string(data)}, " ")
}
