package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VirtualChannelBandwidthPercentageOptions struct {

	// 虚拟通道带宽百分比控制量（%）。取值范围为[0-100]。默认：65。
	VirtualChannelBandwidthPercentageValue *int32 `json:"virtual_channel_bandwidth_percentage_value,omitempty"`
}

func (o VirtualChannelBandwidthPercentageOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VirtualChannelBandwidthPercentageOptions struct{}"
	}

	return strings.Join([]string{"VirtualChannelBandwidthPercentageOptions", string(data)}, " ")
}
