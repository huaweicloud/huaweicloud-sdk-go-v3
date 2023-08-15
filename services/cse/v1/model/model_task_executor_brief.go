package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskExecutorBrief 任务元数据
type TaskExecutorBrief struct {

	// 子任务持续时长
	Duration *int64 `json:"duration,omitempty"`

	// 子任务描述
	Description *string `json:"description,omitempty"`
}

func (o TaskExecutorBrief) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskExecutorBrief struct{}"
	}

	return strings.Join([]string{"TaskExecutorBrief", string(data)}, " ")
}
