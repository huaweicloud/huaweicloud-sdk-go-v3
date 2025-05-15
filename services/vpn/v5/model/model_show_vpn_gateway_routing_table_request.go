package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVpnGatewayRoutingTableRequest Request Object
type ShowVpnGatewayRoutingTableRequest struct {

	// VPN网关实例ID
	VgwId string `json:"vgw_id"`

	// 分页查询时每页返回的记录数量
	Limit *int32 `json:"limit,omitempty"`

	// 上一页最后一条资源记录的创建时间，为空时为查询第一页。使用说明：必须与limit一起使用。
	Marker *string `json:"marker,omitempty"`

	// 分页查询的偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 是否包含下一跳资源信息
	IsIncludeNexthopResource *bool `json:"is_include_nexthop_resource,omitempty"`
}

func (o ShowVpnGatewayRoutingTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVpnGatewayRoutingTableRequest struct{}"
	}

	return strings.Join([]string{"ShowVpnGatewayRoutingTableRequest", string(data)}, " ")
}
