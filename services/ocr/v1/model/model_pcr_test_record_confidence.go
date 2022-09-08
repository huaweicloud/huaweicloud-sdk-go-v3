package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PcrTestRecordConfidence struct {

	// 姓名的置信度
	Name *float32 `json:"name,omitempty"`

	// 核酸检测采样时间的置信度
	SamplingTime *float32 `json:"sampling_time,omitempty"`

	// 核酸检测结果更新时间的置信度
	TestTime *float32 `json:"test_time,omitempty"`

	// 核酸检测结果的置信度
	TestResult *float32 `json:"test_result,omitempty"`
}

func (o PcrTestRecordConfidence) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PcrTestRecordConfidence struct{}"
	}

	return strings.Join([]string{"PcrTestRecordConfidence", string(data)}, " ")
}
