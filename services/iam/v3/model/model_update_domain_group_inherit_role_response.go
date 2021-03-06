package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateDomainGroupInheritRoleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateDomainGroupInheritRoleResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateDomainGroupInheritRoleResponse struct{}"
	}

	return strings.Join([]string{"UpdateDomainGroupInheritRoleResponse", string(data)}, " ")
}
