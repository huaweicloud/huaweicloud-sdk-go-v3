package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 批量启动\\停止任务的详情
type TaskBean struct {
	// 任务ID, 可为空

	TaskId *string `json:"task_id,omitempty"`
}

func (o TaskBean) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskBean struct{}"
	}

	return strings.Join([]string{"TaskBean", string(data)}, " ")
}
