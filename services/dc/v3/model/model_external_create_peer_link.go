package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExternalCreatePeerLink struct {

	// 专线对等连接
	Id *string `json:"id,omitempty"`

	// 租户ID
	TenantId *string `json:"tenant_id,omitempty"`

	// 名称
	Name string `json:"name"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 全域接入网关ID
	GlobalDcGatewayId *string `json:"global_dc_gateway_id,omitempty"`

	BandwidthInfo *BandwidthInfoExternal `json:"bandwidth_info,omitempty"`

	PeerSite *PeerSiteExternal `json:"peer_site"`

	// '状态信息' - ACTIVE 正常 - ERROR 异常
	Status *string `json:"status,omitempty"`

	// 失败原因
	Reason *string `json:"reason,omitempty"`

	// 创建时间
	CreatedTime *sdktime.SdkTime `json:"created_time,omitempty"`

	// 更新时间
	UpdatedTime *sdktime.SdkTime `json:"updated_time,omitempty"`
}

func (o ExternalCreatePeerLink) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExternalCreatePeerLink struct{}"
	}

	return strings.Join([]string{"ExternalCreatePeerLink", string(data)}, " ")
}
