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
	// 请求ID
	RequestId *string `json:"request_id,omitempty"`
	// 安全组列表响应体
	SecurityGroup []SecurityGroup `json:"security_group,omitempty"`
	// 分页信息
	PageInfo *string `json:"page_info,omitempty"`
}

func (o ListSecurityGroupsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListSecurityGroupsResponse", string(data)}, " ")
}
