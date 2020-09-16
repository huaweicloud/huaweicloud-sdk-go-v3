/*
 * VPC
 *
 * VPC Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListSecurityGroupsResponse struct {
	// 安全组列表对象
	SecurityGroups *[]SecurityGroup `json:"security_groups,omitempty"`
}

func (o ListSecurityGroupsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListSecurityGroupsResponse", string(data)}, " ")
}
