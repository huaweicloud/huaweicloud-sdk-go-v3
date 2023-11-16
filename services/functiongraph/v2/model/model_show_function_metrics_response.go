package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFunctionMetricsResponse Response Object
type ShowFunctionMetricsResponse struct {

	// 并发数
	ConcurrencyNum *[]SlaReportsValue `json:"concurrency_num,omitempty"`

	// 调用次数
	Count *[]SlaReportsValue `json:"count,omitempty"`

	// 平均时延，单位毫秒
	Duration *[]SlaReportsValue `json:"duration,omitempty"`

	// 错误次数
	FailCount *[]SlaReportsValue `json:"fail_count,omitempty"`

	// 函数错误次数
	FunctionErrorCount *[]SlaReportsValue `json:"function_error_count,omitempty"`

	// 系统错误次数
	SystemErrorCount *[]SlaReportsValue `json:"system_error_count,omitempty"`

	// 弹性实例数
	InstanceNum *[]SlaReportsValue `json:"instance_num,omitempty"`

	// 最大时延，单位毫秒
	MaxDuration *[]SlaReportsValue `json:"max_duration,omitempty"`

	// 最小时延，单位毫秒
	MinDuration *[]SlaReportsValue `json:"min_duration,omitempty"`

	// 被拒绝次数
	RejectCount *[]SlaReportsValue `json:"reject_count,omitempty"`

	// 预留实例数
	ReservedInstanceNum *[]SlaReportsValue `json:"reserved_instance_num,omitempty"`
	HttpStatusCode      int                `json:"-"`
}

func (o ShowFunctionMetricsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFunctionMetricsResponse struct{}"
	}

	return strings.Join([]string{"ShowFunctionMetricsResponse", string(data)}, " ")
}
