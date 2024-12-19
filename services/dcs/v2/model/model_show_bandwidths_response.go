package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBandwidthsResponse Response Object
type ShowBandwidthsResponse struct {

	// 实例当前带宽(Mbit/s)。
	Bandwidth *int32 `json:"bandwidth,omitempty"`

	// 实例最大带宽(Mbit/s)。
	MaxBandwidth *int32 `json:"max_bandwidth,omitempty"`

	// 是否支持调带宽。
	AllowModify *bool `json:"allow_modify,omitempty"`

	// 分片带宽列表。
	GroupBandwidths *[]GroupBandwidthInfo `json:"group_bandwidths,omitempty"`
	HttpStatusCode  int                   `json:"-"`
}

func (o ShowBandwidthsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBandwidthsResponse struct{}"
	}

	return strings.Join([]string{"ShowBandwidthsResponse", string(data)}, " ")
}
