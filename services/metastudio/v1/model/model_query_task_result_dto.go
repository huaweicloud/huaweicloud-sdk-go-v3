package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryTaskResultDto 查询任务管理的返回结果DTO对象
type QueryTaskResultDto struct {

	// 资源
	Resource *string `json:"resource,omitempty"`

	// 任务名称
	TaskName *string `json:"task_name,omitempty"`

	// 任务开始时间
	BeginTime *string `json:"begin_time,omitempty"`

	// 任务结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 任务类型
	TaskType *string `json:"task_type,omitempty"`

	// 操作人
	Operator *string `json:"operator,omitempty"`

	// 导入总数
	TotalCount *int32 `json:"total_count,omitempty"`

	// 成功数
	SuccessCount *int32 `json:"success_count,omitempty"`

	// 失败数
	FailureCount *int32 `json:"failure_count,omitempty"`

	// 任务状态
	TaskStatus *string `json:"task_status,omitempty"`

	// 任务ID
	TaskId *string `json:"task_id,omitempty"`

	// 操作用户ID
	UserId *string `json:"user_id,omitempty"`
}

func (o QueryTaskResultDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryTaskResultDto struct{}"
	}

	return strings.Join([]string{"QueryTaskResultDto", string(data)}, " ")
}
