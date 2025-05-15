package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SiteConnection 分支连接。
type SiteConnection struct {

	// 实例ID。
	Id string `json:"id"`

	// 分支网络ID。
	SiteNetworkId string `json:"site_network_id"`

	State *SiteConnectionStateEnum `json:"state"`

	// 实例创建时间。UTC时间格式，yyyy-MM-ddTHH:mm:ss。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 实例更新时间。UTC时间格式，yyyy-MM-ddTHH:mm:ss。
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	// 分支网络连接的两个端点定义，长度固定为2的数组。
	EdgePair []DirectedEdge `json:"edge_pair"`

	CrossRegionType *CrossRegionTypeEnum `json:"cross_region_type"`

	// 全域互联带宽ID。
	GlobalConnectionBandwidthId *string `json:"global_connection_bandwidth_id,omitempty"`

	// 带宽值，单位Mbps。
	BandwidthSize *int64 `json:"bandwidth_size,omitempty"`

	// 是否冻结。
	IsFrozen bool `json:"is_frozen"`

	FrozenEffect *FrozenEffectEnum `json:"frozen_effect,omitempty"`

	// 是否绑定带宽包。
	IsBindBandwidth *bool `json:"is_bind_bandwidth,omitempty"`
}

func (o SiteConnection) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SiteConnection struct{}"
	}

	return strings.Join([]string{"SiteConnection", string(data)}, " ")
}
