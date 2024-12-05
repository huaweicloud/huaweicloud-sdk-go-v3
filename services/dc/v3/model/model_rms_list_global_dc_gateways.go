package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RmsListGlobalDcGateways struct {

	// 唯一ID
	Id *string `json:"id,omitempty"`

	// 租户ID
	TenantId *string `json:"tenant_id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 全球中心网路ID
	GlobalCenterNetworkId *string `json:"global_center_network_id,omitempty"`

	// 站点网络ID
	SiteNetworkId *string `json:"site_network_id,omitempty"`

	// 云连接ID
	CloudConnectionId *string `json:"cloud_connection_id,omitempty"`

	// BGP模式AS号
	BgpAsn *int64 `json:"bgp_asn,omitempty"`

	// 局点ID
	RegionId *string `json:"region_id,omitempty"`

	// 位置名称
	LocationName *string `json:"location_name,omitempty"`

	Locales *Locale `json:"locales,omitempty"`

	// 当前对等体数量
	CurrentPeerLinkCount *int32 `json:"current_peer_link_count,omitempty"`

	// 有效对等体数量
	AvailablePeerLinkCount *int32 `json:"available_peer_link_count,omitempty"`

	// 是否冻结：true-是，false-否
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// 状态：ACTIVE-正常，ERROR-异常
	Status *string `json:"status,omitempty"`

	// 创建时间
	CreatedTime *sdktime.SdkTime `json:"created_time,omitempty"`

	// 更新时间
	UpdatedTime *sdktime.SdkTime `json:"updated_time,omitempty"`

	// 返回给RMS的资源标签
	Tags map[string]string `json:"tags,omitempty"`
}

func (o RmsListGlobalDcGateways) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RmsListGlobalDcGateways struct{}"
	}

	return strings.Join([]string{"RmsListGlobalDcGateways", string(data)}, " ")
}
