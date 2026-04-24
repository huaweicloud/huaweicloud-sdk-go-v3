package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAiAssistantUsageFrequencyRequest Request Object
type ShowAiAssistantUsageFrequencyRequest struct {

	// 返回统计个数，默认为前10个用户的统计。
	Count *int32 `json:"count,omitempty"`

	// 开始时间：由日期加时间组成，UTC格式，例如“2021-05-11T11:45:42Z”。
	StartTime string `json:"start_time"`

	// 结束时间：由日期加时间组成，UTC格式，例如“2021-05-11T11:45:42Z”。
	EndTime string `json:"end_time"`
}

func (o ShowAiAssistantUsageFrequencyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAiAssistantUsageFrequencyRequest struct{}"
	}

	return strings.Join([]string{"ShowAiAssistantUsageFrequencyRequest", string(data)}, " ")
}
