package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskSuccessRate 单个应用的部署成功率
type TaskSuccessRate struct {

	// 任务id
	TaskId *string `json:"task_id,omitempty"`

	// 应用名称
	TaskName *string `json:"task_name,omitempty"`

	// 成功率
	SuccessRate *string `json:"success_rate,omitempty"`

	// 部署记录数
	RecordCount *int32 `json:"record_count,omitempty"`

	// 成功的部署记录数
	SuccessRecordCount *int32 `json:"success_record_count,omitempty"`
}

func (o TaskSuccessRate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskSuccessRate struct{}"
	}

	return strings.Join([]string{"TaskSuccessRate", string(data)}, " ")
}
