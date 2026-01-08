package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TranscodeInfoResult struct {

	// 转码模板名
	TemplateName *string `json:"template_name,omitempty"`

	// 转码进度
	Progress *int32 `json:"progress,omitempty"`

	// 转码开始处理时间，格式按照RFC3339，UTC时间，如2020-09-01T18:50:20Z
	StartTime *string `json:"start_time,omitempty"`

	// 等待时长，单位为秒，当值为非-1时有效
	WaitingTime *int32 `json:"waiting_time,omitempty"`

	// 处理时长，单位为秒，当值为非-1时有效
	ProcessTime *int32 `json:"process_time,omitempty"`
}

func (o TranscodeInfoResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TranscodeInfoResult struct{}"
	}

	return strings.Join([]string{"TranscodeInfoResult", string(data)}, " ")
}
