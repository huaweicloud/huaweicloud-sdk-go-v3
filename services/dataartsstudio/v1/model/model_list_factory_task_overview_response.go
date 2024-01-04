package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFactoryTaskOverviewResponse Response Object
type ListFactoryTaskOverviewResponse struct {

	// 失败的实例数量
	FailCount *int32 `json:"fail_count,omitempty"`

	// 强制成功的实例数量
	ForceSuccessCount *int32 `json:"force_success_count,omitempty"`

	// 冻结的实例数量
	FreezeCount *int32 `json:"freeze_count,omitempty"`

	// 忽略失败的实例数量
	IgnoreSuccessCount *int32 `json:"ignore_success_count,omitempty"`

	// 取消的实例数量
	ManualStopCount *int32 `json:"manual_stop_count,omitempty"`

	// 暂时的实例数量
	PauseCount *int32 `json:"pause_count,omitempty"`

	// 运行中的实例数量
	RunningCount *int32 `json:"running_count,omitempty"`

	// 异常的实例数量
	RunningExceptionCount *int32 `json:"running_exception_count,omitempty"`

	// 跳过的实例数量
	SkipCount *int32 `json:"skip_count,omitempty"`

	// 运行成功的实例数量
	SuccessCount *int32 `json:"success_count,omitempty"`

	// 实例总数
	TotalCount *int32 `json:"total_count,omitempty"`

	// 待确认执行的实例数量
	WaitingConfirmCount *int32 `json:"waiting_confirm_count,omitempty"`

	// 等待运行的实例数量
	WaitingCount   *int32 `json:"waiting_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListFactoryTaskOverviewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFactoryTaskOverviewResponse struct{}"
	}

	return strings.Join([]string{"ListFactoryTaskOverviewResponse", string(data)}, " ")
}
