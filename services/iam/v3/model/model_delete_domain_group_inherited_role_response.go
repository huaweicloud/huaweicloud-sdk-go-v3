package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteDomainGroupInheritedRoleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDomainGroupInheritedRoleResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteDomainGroupInheritedRoleResponse struct{}"
	}

	return strings.Join([]string{"DeleteDomainGroupInheritedRoleResponse", string(data)}, " ")
}
