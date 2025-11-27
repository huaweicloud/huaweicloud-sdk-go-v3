package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPrivateSnatsRequest Request Object
type ListPrivateSnatsRequest struct {

	// 功能说明：每页返回的个数。 取值范围：1~2000。 默认值：2000。
	Limit *int32 `json:"limit,omitempty"`

	// 功能说明：分页查询起始的资源ID，为空时查询第一页。 值从上一次查询的PageInfo中的next_marker或者previous_marker中获取。
	Marker *string `json:"marker,omitempty"`

	// 是否查询前一页。
	PageReverse *bool `json:"page_reverse,omitempty"`

	// SNAT规则的ID。
	Id *[]string `json:"id,omitempty"`

	// SNAT规则的描述。长度范围小于等于255个字符，不能包含“<”和“>”。
	Description *[]string `json:"description,omitempty"`

	// 私网NAT网关实例的ID。
	GatewayId *[]string `json:"gateway_id,omitempty"`

	// 规则匹配的CIDR。
	Cidr *[]string `json:"cidr,omitempty"`

	// 规则匹配的子网的ID。
	VirsubnetId *[]string `json:"virsubnet_id,omitempty"`

	// 中转IP的ID。
	TransitIpId *[]string `json:"transit_ip_id,omitempty"`

	// 中转IP地址。
	TransitIpAddress *[]string `json:"transit_ip_address,omitempty"`

	// 企业项目ID。创建SNAT规则时，关联的企业项目ID。
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`
}

func (o ListPrivateSnatsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrivateSnatsRequest struct{}"
	}

	return strings.Join([]string{"ListPrivateSnatsRequest", string(data)}, " ")
}
