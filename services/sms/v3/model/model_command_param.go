package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CommandParam 命令响应参数
type CommandParam struct {

	// 任务ID
	TaskId *string `json:"task_id,omitempty"`

	// 桶名
	Bucket *string `json:"bucket,omitempty"`
}

func (o CommandParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommandParam struct{}"
	}

	return strings.Join([]string{"CommandParam", string(data)}, " ")
}
