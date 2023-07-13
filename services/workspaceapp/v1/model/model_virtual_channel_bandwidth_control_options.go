package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VirtualChannelBandwidthControlOptions struct {

	// 虚拟通道带宽控制量（Kbps）。取值范围为[500-20000]。默认：20000。
	VirtualChannelBandwidthControlValue *int32 `json:"virtual_channel_bandwidth_control_value,omitempty"`
}

func (o VirtualChannelBandwidthControlOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VirtualChannelBandwidthControlOptions struct{}"
	}

	return strings.Join([]string{"VirtualChannelBandwidthControlOptions", string(data)}, " ")
}
