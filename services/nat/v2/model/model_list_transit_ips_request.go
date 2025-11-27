package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTransitIpsRequest Request Object
type ListTransitIpsRequest struct {

	// 功能说明：每页返回的个数。 取值范围：1~2000。 默认值：2000。
	Limit *int32 `json:"limit,omitempty"`

	// 功能说明：分页查询起始的资源ID，为空时查询第一页。 值从上一次查询的PageInfo中的next_marker或者previous_marker中获取。
	Marker *string `json:"marker,omitempty"`

	// 是否查询前一页。
	PageReverse *bool `json:"page_reverse,omitempty"`

	// 中转IP的ID。
	Id *[]string `json:"id,omitempty"`

	// 中转IP的网络接口ID。
	NetworkInterfaceId *[]string `json:"network_interface_id,omitempty"`

	// 中转IP地址。
	IpAddress *[]string `json:"ip_address,omitempty"`

	// 中转IP绑定的私网NAT网关实例的ID。
	GatewayId *[]string `json:"gateway_id,omitempty"`

	// 企业项目ID。创建中转IP时，关联的企业项目ID。
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`

	// 当前租户子网的ID。
	VirsubnetId *[]string `json:"virsubnet_id,omitempty"`

	// 中转子网的ID。
	TransitSubnetId *[]string `json:"transit_subnet_id,omitempty"`
}

func (o ListTransitIpsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTransitIpsRequest struct{}"
	}

	return strings.Join([]string{"ListTransitIpsRequest", string(data)}, " ")
}
