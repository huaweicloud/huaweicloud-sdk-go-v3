package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListPrivilegesResponse struct {
	// 是否有权限

	HasPrivilege   *int32 `json:"has_privilege,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListPrivilegesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListPrivilegesResponse struct{}"
	}

	return strings.Join([]string{"ListPrivilegesResponse", string(data)}, " ")
}
