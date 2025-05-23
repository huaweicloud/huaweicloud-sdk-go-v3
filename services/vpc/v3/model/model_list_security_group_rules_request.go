package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityGroupRulesRequest Request Object
type ListSecurityGroupRulesRequest struct {

	// 功能说明：每页返回个数 取值范围：0-2000
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询起始的资源ID，为空时查询第一页
	Marker *string `json:"marker,omitempty"`

	// 功能说明：安全组规则ID，支持多个ID过滤
	Id *[]string `json:"id,omitempty"`

	// 功能说明：安全组规则所属安全组ID，支持多个ID过滤
	SecurityGroupId *[]string `json:"security_group_id,omitempty"`

	// 功能说明：安全组规则协议，支持多条过滤
	Protocol *[]string `json:"protocol,omitempty"`

	// 功能说明：安全组规则的描述，支持多个描述同时过滤
	Description *[]string `json:"description,omitempty"`

	// 功能说明：远端安全组ID，支持多ID过滤
	RemoteGroupId *[]string `json:"remote_group_id,omitempty"`

	// 功能说明：安全组规则方向
	Direction *string `json:"direction,omitempty"`

	// 功能说明：安全组规则生效策略。 取值范围：allow表示允许，deny表示拒绝。
	Action *string `json:"action,omitempty"`

	// 功能说明：远端IP地址 取值范围：cidr格式
	RemoteIpPrefix *string `json:"remote_ip_prefix,omitempty"`

	// 功能说明：优先级，支持多条过滤。
	Priority *[]int32 `json:"priority,omitempty"`

	// 功能说明：IP协议类型，支持多条过滤。 取值范围：IPv4,IPv6,ipv4,ipv6
	Ethertype *[]string `json:"ethertype,omitempty"`

	// 功能说明：远端IP地址组ID，支持多ID过滤。
	RemoteAddressGroupId *[]string `json:"remote_address_group_id,omitempty"`

	// 功能说明：是否启用安全组规则，不支持多值过滤。 取值范围：true, false。
	Enabled *bool `json:"enabled,omitempty"`
}

func (o ListSecurityGroupRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityGroupRulesRequest struct{}"
	}

	return strings.Join([]string{"ListSecurityGroupRulesRequest", string(data)}, " ")
}
