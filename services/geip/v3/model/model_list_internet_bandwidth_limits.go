package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListInternetBandwidthLimits struct {

	// 全域公网带宽限制的ID
	Id *string `json:"id,omitempty"`

	// 全域公网带宽的计费模式
	ChargeMode *string `json:"charge_mode,omitempty"`

	// 该类型全域公网带宽可购买的最小size
	MinSize *int32 `json:"min_size,omitempty"`

	ExtLimit *ExtLimitPojo `json:"ext_limit,omitempty"`

	// 该类型全域公网带宽可购买的最大size
	MaxSize *int32 `json:"max_size,omitempty"`

	// 全域公网带宽类型
	Type *string `json:"type,omitempty"`
}

func (o ListInternetBandwidthLimits) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInternetBandwidthLimits struct{}"
	}

	return strings.Join([]string{"ListInternetBandwidthLimits", string(data)}, " ")
}
