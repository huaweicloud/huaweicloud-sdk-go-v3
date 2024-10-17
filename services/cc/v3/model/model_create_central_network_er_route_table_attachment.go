package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCentralNetworkErRouteTableAttachment 创建中心网络ER附件的请求体。
type CreateCentralNetworkErRouteTableAttachment struct {

	// 实例名字。
	Name string `json:"name"`

	// 实例描述。不支持 <>。
	Description *string `json:"description,omitempty"`

	// 企业路由器的ID。
	EnterpriseRouterId string `json:"enterprise_router_id"`

	// 企业路由器的项目ID。
	EnterpriseRouterProjectId string `json:"enterprise_router_project_id"`

	// ER路由器的regionID。
	EnterpriseRouterRegionId string `json:"enterprise_router_region_id"`

	// 中心网络平面ID。
	CentralNetworkPlaneId string `json:"central_network_plane_id"`

	// 中心网络附件对端实例的连接ID，企业路由器的连接ID或者GDGW的连接ID。
	AttachmentId *string `json:"attachment_id,omitempty"`

	// 企业路由器的路由表ID。
	EnterpriseRouterTableId string `json:"enterprise_router_table_id"`

	// 被挂载的企业路由器的项目ID。
	AttachedErTableProjectId string `json:"attached_er_table_project_id"`

	// ER路由器的regionID。
	AttachedErTableRegionId string `json:"attached_er_table_region_id"`

	// 被挂载的企业路由器ID。
	AttachedErId string `json:"attached_er_id"`

	// 被挂载的企业路由器的路由表ID。
	AttachedErTableId string `json:"attached_er_table_id"`

	HostedCloud *HostedCloudEnum `json:"hosted_cloud"`
}

func (o CreateCentralNetworkErRouteTableAttachment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCentralNetworkErRouteTableAttachment struct{}"
	}

	return strings.Join([]string{"CreateCentralNetworkErRouteTableAttachment", string(data)}, " ")
}
