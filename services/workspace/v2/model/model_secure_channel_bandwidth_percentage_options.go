package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SecureChannelBandwidthPercentageOptions struct {

	// 安全通道带宽百分比控制量（%）。取值范围为[0-100]。默认：30。
	SecureChannelBandwidthPercentageValue *int32 `json:"secure_channel_bandwidth_percentage_value,omitempty"`
}

func (o SecureChannelBandwidthPercentageOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecureChannelBandwidthPercentageOptions struct{}"
	}

	return strings.Join([]string{"SecureChannelBandwidthPercentageOptions", string(data)}, " ")
}
