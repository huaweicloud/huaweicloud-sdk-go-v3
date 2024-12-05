package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExternalShowPeerLink 对外对等体对象
type ExternalShowPeerLink struct {

	// 唯一ID
	Id *string `json:"id,omitempty"`

	// 租户ID
	TenantId *string `json:"tenant_id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 全球接入网关ID
	GlobalDcGatewayId *string `json:"global_dc_gateway_id,omitempty"`

	// 创建者
	CreateOwner *string `json:"create_owner,omitempty"`

	BandwidthInfo *BandwidthInfoExternal `json:"bandwidth_info,omitempty"`

	PeerSite *PeerSiteExternal `json:"peer_site,omitempty"`

	// 状态：ACTIVE-正常，ERROR-异常
	Status *string `json:"status,omitempty"`

	// 原因
	Reason *string `json:"reason,omitempty"`

	// 创建时间
	CreatedTime *sdktime.SdkTime `json:"created_time,omitempty"`

	// 修改时间
	UpdatedTime *sdktime.SdkTime `json:"updated_time,omitempty"`
}

func (o ExternalShowPeerLink) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExternalShowPeerLink struct{}"
	}

	return strings.Join([]string{"ExternalShowPeerLink", string(data)}, " ")
}
