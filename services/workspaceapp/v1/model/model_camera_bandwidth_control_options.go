package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CameraBandwidthControlOptions struct {

	// 摄像头带宽控制量（Kbps）。取值范围为[0-10000]。默认：10000。
	CameraBandwidthControlValue *int32 `json:"camera_bandwidth_control_value,omitempty"`
}

func (o CameraBandwidthControlOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CameraBandwidthControlOptions struct{}"
	}

	return strings.Join([]string{"CameraBandwidthControlOptions", string(data)}, " ")
}
