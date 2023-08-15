package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskData struct {

	// 作业流详情
	Streams []TaskStream `json:"streams"`

	// 作业id
	TaskId *string `json:"task_id,omitempty"`

	StartTimeInfo *StartTimeInfo `json:"start_time_info,omitempty"`

	SourceUsageEstimate *TaskSourceUsageEstimate `json:"source_usage_estimate,omitempty"`
}

func (o TaskData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskData struct{}"
	}

	return strings.Join([]string{"TaskData", string(data)}, " ")
}
