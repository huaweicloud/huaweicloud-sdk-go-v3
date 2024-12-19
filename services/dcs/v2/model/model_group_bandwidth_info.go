package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GroupBandwidthInfo struct {

	// 分片ID。
	GroupId *string `json:"group_id,omitempty"`

	// 更新时间，UTC时间。
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 当前带宽(Mbit/s)。
	Bandwidth *int32 `json:"bandwidth,omitempty"`

	// 最大带宽(Mbit/s)。
	MaxBandwidth *int32 `json:"max_bandwidth,omitempty"`

	// 基准带宽(Mbit/s)。
	AssuredBandwidth *int32 `json:"assured_bandwidth,omitempty"`
}

func (o GroupBandwidthInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupBandwidthInfo struct{}"
	}

	return strings.Join([]string{"GroupBandwidthInfo", string(data)}, " ")
}
