package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CentralNetworkAttachment 中心网络附件详情。
type CentralNetworkAttachment struct {

	// 资源ID标识符。
	Id string `json:"id"`

	// 实例名字。
	Name string `json:"name"`

	// 实例描述。不支持 <>。
	Description *string `json:"description,omitempty"`

	// 实例所属帐号ID。
	DomainId string `json:"domain_id"`

	State *CentralNetworkConnectionStateEnum `json:"state"`

	// 实例创建时间。UTC时间格式，yyyy-MM-ddTHH:mm:ss。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 实例更新时间。UTC时间格式，yyyy-MM-ddTHH:mm:ss。
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	// 资源ID标识符。
	CentralNetworkId string `json:"central_network_id"`

	// 资源ID标识符。
	CentralNetworkPlaneId string `json:"central_network_plane_id"`

	// 资源ID标识符。
	GlobalConnectionBandwidthId *string `json:"global_connection_bandwidth_id,omitempty"`

	BandwidthType *BandwidthTypeEnum `json:"bandwidth_type"`

	// 带宽值定义，单位Mbps。
	BandwidthSize *int64 `json:"bandwidth_size,omitempty"`

	// 是否冻结
	IsFrozen bool `json:"is_frozen"`

	// 资源ID标识符。
	EnterpriseRouterId string `json:"enterprise_router_id"`

	// 实例所属项目ID。
	EnterpriseRouterProjectId string `json:"enterprise_router_project_id"`

	// RegionID。
	EnterpriseRouterRegionId string `json:"enterprise_router_region_id"`

	// 资源ID标识符。
	EnterpriseRouterAttachmentId *string `json:"enterprise_router_attachment_id,omitempty"`

	AttachmentInstanceType *AttachmentInstanceTypeEnum `json:"attachment_instance_type"`

	// 资源ID标识符。
	AttachmentInstanceId string `json:"attachment_instance_id"`

	// 资源ID标识符。
	AttachmentId *string `json:"attachment_id,omitempty"`

	// 实例所属项目ID。
	AttachmentInstanceProjectId string `json:"attachment_instance_project_id"`

	// RegionID。
	AttachmentInstanceRegionId string `json:"attachment_instance_region_id"`

	// 站点编码定义
	AttachmentInstanceSiteCode string `json:"attachment_instance_site_code"`

	// 站点编码定义
	EnterpriseRouterSiteCode string `json:"enterprise_router_site_code"`

	SpecificationValue *CentralNetworkAttachmentSpecificationValueInfo `json:"specification_value,omitempty"`
}

func (o CentralNetworkAttachment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CentralNetworkAttachment struct{}"
	}

	return strings.Join([]string{"CentralNetworkAttachment", string(data)}, " ")
}
