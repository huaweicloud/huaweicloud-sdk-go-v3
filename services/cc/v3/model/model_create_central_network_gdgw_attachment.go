package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateCentralNetworkGdgwAttachment struct {

	// 实例名字。
	Name string `json:"name"`

	// 实例描述。不支持 <>。
	Description *string `json:"description,omitempty"`

	// 资源ID标识符。
	EnterpriseRouterId string `json:"enterprise_router_id"`

	// 资源ID标识符。
	GlobalDcGatewayId string `json:"global_dc_gateway_id"`

	// 实例所属项目ID。
	EnterpriseRouterProjectId string `json:"enterprise_router_project_id"`

	// RegionID。
	EnterpriseRouterRegionId string `json:"enterprise_router_region_id"`

	// 实例所属项目ID。
	GlobalDcGatewayProjectId string `json:"global_dc_gateway_project_id"`

	// RegionID。
	GlobalDcGatewayRegionId string `json:"global_dc_gateway_region_id"`

	// 资源ID标识符。
	CentralNetworkPlaneId *string `json:"central_network_plane_id,omitempty"`
}

func (o CreateCentralNetworkGdgwAttachment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCentralNetworkGdgwAttachment struct{}"
	}

	return strings.Join([]string{"CreateCentralNetworkGdgwAttachment", string(data)}, " ")
}
