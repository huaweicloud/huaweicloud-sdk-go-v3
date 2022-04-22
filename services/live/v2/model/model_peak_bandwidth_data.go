package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PeakBandwidthData struct {

	// 带宽峰值，单位为bps。
	Value *int64 `json:"value,omitempty"`

	// 播放域名。
	Domain *string `json:"domain,omitempty"`
}

func (o PeakBandwidthData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PeakBandwidthData struct{}"
	}

	return strings.Join([]string{"PeakBandwidthData", string(data)}, " ")
}
