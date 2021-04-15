package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type NeutronListSecurityGroupsResponse struct {
	// 安全组对象列表

	SecurityGroups *[]NeutronSecurityGroup `json:"security_groups,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o NeutronListSecurityGroupsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "NeutronListSecurityGroupsResponse struct{}"
	}

	return strings.Join([]string{"NeutronListSecurityGroupsResponse", string(data)}, " ")
}
