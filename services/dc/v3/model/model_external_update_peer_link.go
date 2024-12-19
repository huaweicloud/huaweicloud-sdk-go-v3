package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExternalUpdatePeerLink 更新的请求体信息
type ExternalUpdatePeerLink struct {

	// 关联连接的ID
	Id *string `json:"id,omitempty"`

	// 关联连接归属的租户ID
	TenantId *string `json:"tenant_id,omitempty"`

	// 关联连接的名字
	Name *string `json:"name,omitempty"`

	// 关联连接的描述信息
	Description *string `json:"description,omitempty"`

	// 关联连接归属的接入网关ID
	GlobalDcGatewayId *string `json:"global_dc_gateway_id,omitempty"`

	BandwidthInfo *BandwidthInfoExternal `json:"bandwidth_info,omitempty"`

	PeerSite *PeerSiteExternal `json:"peer_site,omitempty"`

	// 关联连接的状态
	Status *string `json:"status,omitempty"`

	// 失败原因
	Reason *string `json:"reason,omitempty"`

	// 创建时间
	CreatedTime *sdktime.SdkTime `json:"created_time,omitempty"`

	// 更新时间
	UpdatedTime *sdktime.SdkTime `json:"updated_time,omitempty"`
}

func (o ExternalUpdatePeerLink) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExternalUpdatePeerLink struct{}"
	}

	return strings.Join([]string{"ExternalUpdatePeerLink", string(data)}, " ")
}
