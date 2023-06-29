package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ClipboardBandwidthControlOptions struct {

	// 剪切板带宽控制量（Kbps）。取值范围为[100-2000]。默认：1000。
	ClipboardBandwidthControlValue *int32 `json:"clipboard_bandwidth_control_value,omitempty"`
}

func (o ClipboardBandwidthControlOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClipboardBandwidthControlOptions struct{}"
	}

	return strings.Join([]string{"ClipboardBandwidthControlOptions", string(data)}, " ")
}
