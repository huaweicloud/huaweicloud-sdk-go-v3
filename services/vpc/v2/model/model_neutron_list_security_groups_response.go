package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type NeutronListSecurityGroupsResponse struct {
	// 安全组对象列表

	SecurityGroups *[]NeutronSecurityGroup `json:"security_groups,omitempty"`
	// 分页信息

	SecurityGroupsLinks *[]NeutronPageLink `json:"security_groups_links,omitempty"`
	HttpStatusCode      int                `json:"-"`
}

func (o NeutronListSecurityGroupsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "NeutronListSecurityGroupsResponse struct{}"
	}

	return strings.Join([]string{"NeutronListSecurityGroupsResponse", string(data)}, " ")
}
