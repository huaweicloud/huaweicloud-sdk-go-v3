package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetSubJobDetail 子任务详情
type GetSubJobDetail struct {

	// 子任务的状态。success：成功。running：运行中。failed：失败。waiting：等待执行。
	Status *string `json:"status,omitempty"`

	// 子任务的ID。
	JobId *string `json:"job_id,omitempty"`

	// 子任务的类型。
	JobType *string `json:"job_type,omitempty"`

	// 子任务开始时间。UTC时间，格式：'2016-01-02 15:04:05'
	BeginTime *string `json:"begin_time,omitempty"`

	// 子任务结束时间。UTC时间，格式：'2016-01-02 15:04:05'
	EndTime *string `json:"end_time,omitempty"`

	// 子任务执行失败时的错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 子任务执行失败时的错误原因
	FailReason *string `json:"fail_reason,omitempty"`
}

func (o GetSubJobDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetSubJobDetail struct{}"
	}

	return strings.Join([]string{"GetSubJobDetail", string(data)}, " ")
}
