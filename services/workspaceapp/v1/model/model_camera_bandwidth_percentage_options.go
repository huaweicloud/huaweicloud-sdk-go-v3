package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CameraBandwidthPercentageOptions struct {

	// 摄像头带宽百分比控制量（%）。取值范围为[0-100]。默认：30。
	CameraBandwidthPercentageValue *int32 `json:"camera_bandwidth_percentage_value,omitempty"`
}

func (o CameraBandwidthPercentageOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CameraBandwidthPercentageOptions struct{}"
	}

	return strings.Join([]string{"CameraBandwidthPercentageOptions", string(data)}, " ")
}
