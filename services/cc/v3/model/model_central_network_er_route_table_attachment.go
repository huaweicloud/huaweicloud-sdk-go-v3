package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CentralNetworkErRouteTableAttachment 企业路由器的路由表附件详情。
type CentralNetworkErRouteTableAttachment struct {

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

	// 是否冻结
	IsFrozen bool `json:"is_frozen"`

	BandwidthType *BandwidthTypeEnum `json:"bandwidth_type"`

	// 带宽值，单位Mbps。
	BandwidthSize *int64 `json:"bandwidth_size,omitempty"`

	// 企业路由器的ID。
	EnterpriseRouterId string `json:"enterprise_router_id"`

	// 企业路由器的项目ID。
	EnterpriseRouterProjectId string `json:"enterprise_router_project_id"`

	// ER路由器的regionID。
	EnterpriseRouterRegionId string `json:"enterprise_router_region_id"`

	// 企业路由器的连接ID。
	EnterpriseRouterAttachmentId *string `json:"enterprise_router_attachment_id,omitempty"`

	// 企业路由器的路由表ID。
	EnterpriseRouterTableId string `json:"enterprise_router_table_id"`

	// 中心网络企业路由器的站点编码。
	EnterpriseRouterSiteCode string `json:"enterprise_router_site_code"`

	// 被挂载的企业路由器的项目ID。
	AttachedErTableProjectId string `json:"attached_er_table_project_id"`

	// ER路由器的regionID。
	AttachedErTableRegionId string `json:"attached_er_table_region_id"`

	// 被挂载的企业路由器ID。
	AttachedErId string `json:"attached_er_id"`

	// 被挂载的企业路由器的路由表ID。
	AttachedErTableId string `json:"attached_er_table_id"`

	// 被挂载的企业路由器的连接ID。
	AttachedErAttachmentId *string `json:"attached_er_attachment_id,omitempty"`

	// 被挂载的企业路由器的站点编码。
	AttachedErTableSiteCode string `json:"attached_er_table_site_code"`

	ApprovedState *ApprovedStateEnum `json:"approved_state"`

	HostedCloud *HostedCloudEnum `json:"hosted_cloud,omitempty"`

	// 审批拒绝创建企业路由表附件的原因。
	Reason *string `json:"reason,omitempty"`
}

func (o CentralNetworkErRouteTableAttachment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CentralNetworkErRouteTableAttachment struct{}"
	}

	return strings.Join([]string{"CentralNetworkErRouteTableAttachment", string(data)}, " ")
}
