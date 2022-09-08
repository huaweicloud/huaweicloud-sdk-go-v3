package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 单个任务的执行成功率
type TaskSuccessRate struct {

	// 任务id
	TaskId *string `json:"task_id,omitempty"`

	// 任务名称
	TaskName *string `json:"task_name,omitempty"`

	// 成功率
	SuccessRate *string `json:"success_rate,omitempty"`

	// 执行记录数
	RecordCount *int32 `json:"record_count,omitempty"`

	// 成功的执行记录数
	SuccessRecordCount *int32 `json:"success_record_count,omitempty"`
}

func (o TaskSuccessRate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskSuccessRate struct{}"
	}

	return strings.Join([]string{"TaskSuccessRate", string(data)}, " ")
}
