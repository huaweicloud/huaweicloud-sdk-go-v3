package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type NeutronListFirewallPoliciesRequest struct {

	// 每页返回的个数
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 分页查询起始的资源ID，为空时查询第一页
	Marker *string `json:"marker,omitempty" xml:"marker"`

	// 使用网络ACL策略ID过滤网络ACL策略
	Id *[]string `json:"id,omitempty" xml:"id"`

	// 使用name过滤网络ACL策略
	Name *[]string `json:"name,omitempty" xml:"name"`

	// 使用网络ACL策略描述过滤查询网络ACL策略
	Description *[]string `json:"description,omitempty" xml:"description"`

	// 使用tenant_id过滤查询网络ACL策略
	TenantId *string `json:"tenant_id,omitempty" xml:"tenant_id"`
}

func (o NeutronListFirewallPoliciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronListFirewallPoliciesRequest struct{}"
	}

	return strings.Join([]string{"NeutronListFirewallPoliciesRequest", string(data)}, " ")
}
