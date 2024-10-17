package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CentralNetworkGdgwAttachment 中心GDGW附件详情。
type CentralNetworkGdgwAttachment struct {

	// 实例ID。
	Id string `json:"id"`

	// 实例名字。
	Name string `json:"name"`

	// 实例描述。不支持 <>。
	Description *string `json:"description,omitempty"`

	// 实例所属账号ID。
	DomainId string `json:"domain_id"`

	State *CentralNetworkConnectionStateEnum `json:"state"`

	// 实例创建时间。UTC时间格式，yyyy-MM-ddTHH:mm:ss。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 实例更新时间。UTC时间格式，yyyy-MM-ddTHH:mm:ss。
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	// 中心网络ID。
	CentralNetworkId string `json:"central_network_id"`

	// 中心网络平面ID。
	CentralNetworkPlaneId string `json:"central_network_plane_id"`

	// 全域互联带宽ID。
	GlobalConnectionBandwidthId *string `json:"global_connection_bandwidth_id,omitempty"`

	BandwidthType *BandwidthTypeEnum `json:"bandwidth_type"`

	// 带宽值，单位Mbps。
	BandwidthSize *int64 `json:"bandwidth_size,omitempty"`

	// 是否冻结
	IsFrozen bool `json:"is_frozen"`

	// 企业路由器的ID。
	EnterpriseRouterId string `json:"enterprise_router_id"`

	// 企业路由器的项目ID。
	EnterpriseRouterProjectId string `json:"enterprise_router_project_id"`

	// ER路由器的regionID。
	EnterpriseRouterRegionId string `json:"enterprise_router_region_id"`

	// 企业路由器的连接ID。
	EnterpriseRouterAttachmentId *string `json:"enterprise_router_attachment_id,omitempty"`

	// GDGW的连接ID。
	GlobalDcGatewayPeerLinkId *string `json:"global_dc_gateway_peer_link_id,omitempty"`

	// Gdgw的ID。
	GlobalDcGatewayId string `json:"global_dc_gateway_id"`

	// Gdgw的项目ID。
	GlobalDcGatewayProjectId string `json:"global_dc_gateway_project_id"`

	// Gdgw的RegionID。
	GlobalDcGatewayRegionId string `json:"global_dc_gateway_region_id"`

	// 中心网络企业路由器的站点编码。
	EnterpriseRouterSiteCode string `json:"enterprise_router_site_code"`

	// 站点编码定义
	GlobalDcGatewaySiteCode string `json:"global_dc_gateway_site_code"`
}

func (o CentralNetworkGdgwAttachment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CentralNetworkGdgwAttachment struct{}"
	}

	return strings.Join([]string{"CentralNetworkGdgwAttachment", string(data)}, " ")
}
