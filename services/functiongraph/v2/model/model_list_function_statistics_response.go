package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListFunctionStatisticsResponse struct {

	// 调用次数
	Count *[]SlaReportsValue `json:"count,omitempty" xml:"count"`

	// 平均时延，单位毫秒
	Duration *[]SlaReportsValue `json:"duration,omitempty" xml:"duration"`

	// 错误次数
	FailCount *[]SlaReportsValue `json:"fail_count,omitempty" xml:"fail_count"`

	// 最大时延，单位毫秒
	MaxDuration *[]SlaReportsValue `json:"max_duration,omitempty" xml:"max_duration"`

	// 最小时延，单位毫秒
	MinDuration *[]SlaReportsValue `json:"min_duration,omitempty" xml:"min_duration"`

	// 被拒绝次数
	RejectCount    *[]SlaReportsValue `json:"reject_count,omitempty" xml:"reject_count"`
	HttpStatusCode int                `json:"-"`
}

func (o ListFunctionStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFunctionStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListFunctionStatisticsResponse", string(data)}, " ")
}
