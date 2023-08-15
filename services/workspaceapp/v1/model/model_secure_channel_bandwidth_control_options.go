package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SecureChannelBandwidthControlOptions struct {

	// 安全通道带宽控制量（Kbps）。取值范围为[500-20000]。默认：10000。
	SecureChannelBandwidthControlValue *int32 `json:"secure_channel_bandwidth_control_value,omitempty"`
}

func (o SecureChannelBandwidthControlOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecureChannelBandwidthControlOptions struct{}"
	}

	return strings.Join([]string{"SecureChannelBandwidthControlOptions", string(data)}, " ")
}
