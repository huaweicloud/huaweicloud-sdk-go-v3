package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskListVo 实际的数据类型：单个对象，集合 或 NULL
type TaskListVo struct {

	// 测试任务集合
	Tasks *[]TaskVo `json:"tasks,omitempty"`

	// 正在执行任务数
	RunningCount *int32 `json:"running_count,omitempty"`
}

func (o TaskListVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskListVo struct{}"
	}

	return strings.Join([]string{"TaskListVo", string(data)}, " ")
}
