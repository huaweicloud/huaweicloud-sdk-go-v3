package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowQuotasRequest Request Object
type ShowQuotasRequest struct {

	// 支持过滤的配额类型： - er_instance: 企业路由器实例的配额和使用量 - dc_attachment: 云专线网关连接的配额和使用量 - vpc_attachment: VPC连接的配额和使用量 - vpn_attachment: VPN网关连接的配额和使用量 - peering_attachment：云连接实例连接的配额和使用量 - can_attachment: 智能接入网关连接的配额和使用量 - route_table: 路由表的配额和使用量 - static_route: 静态路由的配额和使用量 - vpc_er: 每个vpc可以接入的企业路由器数量和当前使用量 - flow_log: 每个连接可以创建的流日志数量
	Type *[]string `json:"type,omitempty"`

	ErId *[]string `json:"erId,omitempty"`

	RouteTableId *[]string `json:"routeTableId,omitempty"`

	VpcId *[]string `json:"vpcId,omitempty"`
}

func (o ShowQuotasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQuotasRequest struct{}"
	}

	return strings.Join([]string{"ShowQuotasRequest", string(data)}, " ")
}
