package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCentralNetworkGdgwAttachment 创建中心网络GDGW附件的请求体。
type CreateCentralNetworkGdgwAttachment struct {

	// 实例名字。
	Name string `json:"name"`

	// 实例描述。不支持 <>。
	Description *string `json:"description,omitempty"`

	// 企业路由器的ID。
	EnterpriseRouterId string `json:"enterprise_router_id"`

	// Gdgw的ID。
	GlobalDcGatewayId string `json:"global_dc_gateway_id"`

	// Gdgw的项目ID。
	GlobalDcGatewayProjectId string `json:"global_dc_gateway_project_id"`

	// Gdgw的RegionID。
	GlobalDcGatewayRegionId string `json:"global_dc_gateway_region_id"`

	// 企业路由器的项目ID。
	EnterpriseRouterProjectId string `json:"enterprise_router_project_id"`

	// ER路由器的regionID。
	EnterpriseRouterRegionId string `json:"enterprise_router_region_id"`

	// 中心网络平面ID。
	CentralNetworkPlaneId *string `json:"central_network_plane_id,omitempty"`
}

func (o CreateCentralNetworkGdgwAttachment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCentralNetworkGdgwAttachment struct{}"
	}

	return strings.Join([]string{"CreateCentralNetworkGdgwAttachment", string(data)}, " ")
}
