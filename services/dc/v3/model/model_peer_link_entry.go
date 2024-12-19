package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PeerLinkEntry 实例的peer Link的详情
type PeerLinkEntry struct {

	// peer link ID。
	Id *string `json:"id,omitempty"`

	// 租户项目ID
	TenantId *string `json:"tenant_id,omitempty"`

	// 专线内部连接(peer link)名字
	Name *string `json:"name,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 对应的专线全域接入网关ID
	GlobalDcGatewayId *string `json:"global_dc_gateway_id,omitempty"`

	BandwidthInfo *BandWidthInfo `json:"bandwidth_info,omitempty"`

	PeerSite *PeerSite `json:"peer_site,omitempty"`

	Status *PeerLinkStatus `json:"status,omitempty"`

	// 创建时间。
	CreatedTime *sdktime.SdkTime `json:"created_time,omitempty"`

	// 更新时间。
	UpdatedTime *sdktime.SdkTime `json:"updated_time,omitempty"`

	// 创建归属服务名 - cc 云连接 - dc 云专线
	CreateOwner *sdktime.SdkTime `json:"create_owner,omitempty"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`
}

func (o PeerLinkEntry) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PeerLinkEntry struct{}"
	}

	return strings.Join([]string{"PeerLinkEntry", string(data)}, " ")
}
