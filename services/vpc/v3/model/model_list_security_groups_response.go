package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSecurityGroupsResponse struct {
	// 安全组列表响应体

	SecurityGroups *[]SecurityGroup `json:"security_groups,omitempty"`
	// 请求ID

	RequestId *string `json:"request_id,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListSecurityGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityGroupsResponse", string(data)}, " ")
}
