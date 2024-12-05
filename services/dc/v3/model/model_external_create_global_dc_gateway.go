package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExternalCreateGlobalDcGateway 创建成功返回响应体
type ExternalCreateGlobalDcGateway struct {

	// 唯一ID
	Id *string `json:"id,omitempty"`

	// 租户ID
	TenantId *string `json:"tenant_id,omitempty"`

	// 全球接入网关ID
	Name string `json:"name"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 地址簇
	AddressFamily *string `json:"address_family,omitempty"`

	// 全球中心网络ID
	GlobalCenterNetworkId *string `json:"global_center_network_id,omitempty"`

	// 站点网络ID
	SiteNetworkId *string `json:"site_network_id,omitempty"`

	// 云连接ID
	CloudConnectionId *string `json:"cloud_connection_id,omitempty"`

	// BGP的AS号
	BgpAsn *int64 `json:"bgp_asn,omitempty"`

	// 局点ID
	RegionId *string `json:"region_id,omitempty"`

	// 位置名称
	LocationName *string `json:"location_name,omitempty"`

	// 当前对等体数量
	CurrentPeerLinkCount *int32 `json:"current_peer_link_count,omitempty"`

	// 有效对等体数量
	AvailablePeerLinkCount *int32 `json:"available_peer_link_count,omitempty"`

	// 是否冻结 - true 是 - false 否
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// 状态值 - ACTIVE 正常 - ERROR 异常
	Status *string `json:"status,omitempty"`

	// 创建时间
	CreatedTime *sdktime.SdkTime `json:"created_time,omitempty"`

	// 修改时间
	UpdatedTime *sdktime.SdkTime `json:"updated_time,omitempty"`
}

func (o ExternalCreateGlobalDcGateway) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExternalCreateGlobalDcGateway struct{}"
	}

	return strings.Join([]string{"ExternalCreateGlobalDcGateway", string(data)}, " ")
}
