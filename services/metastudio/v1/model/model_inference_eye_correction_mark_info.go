package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InferenceEyeCorrectionMarkInfo 分身数字人推理预处理眼神矫正信息。
type InferenceEyeCorrectionMarkInfo struct {

	// 选取推理数据预处理眼神矫正起始时间。格式：“HH:MM:SS.mmm”。
	EyeCorrectionStartTime *string `json:"eye_correction_start_time,omitempty"`

	// 选取推理数据预处理眼神矫正结束时间。格式：“HH:MM:SS.mmm”。
	EyeCorrectionEndTime *string `json:"eye_correction_end_time,omitempty"`
}

func (o InferenceEyeCorrectionMarkInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InferenceEyeCorrectionMarkInfo struct{}"
	}

	return strings.Join([]string{"InferenceEyeCorrectionMarkInfo", string(data)}, " ")
}
