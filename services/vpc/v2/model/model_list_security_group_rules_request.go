package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityGroupRulesRequest Request Object
type ListSecurityGroupRulesRequest struct {

	// 功能说明：分页查询起始的资源ID，为空时查询第一页
	Marker *string `json:"marker,omitempty"`

	// 每页返回的个数
	Limit *int32 `json:"limit,omitempty"`

	// 安全组ID
	SecurityGroupId *string `json:"security_group_id,omitempty"`

	// 功能说明：远端IP地址 取值范围：cidr格式
	RemoteIpPrefix *string `json:"remote_ip_prefix,omitempty"`
}

func (o ListSecurityGroupRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityGroupRulesRequest struct{}"
	}

	return strings.Join([]string{"ListSecurityGroupRulesRequest", string(data)}, " ")
}
