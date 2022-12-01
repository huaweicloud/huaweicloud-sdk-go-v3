package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 任务项详情
type Task struct {

	// 任务项ID
	TaskId *string `json:"task_id,omitempty"`

	// 任务项名称
	TaskName *string `json:"task_name,omitempty"`

	// 任务项状态
	Status *string `json:"status,omitempty"`

	// 任务项失败的原因
	Reason *string `json:"reason,omitempty"`

	// 创建时间戳
	CreatedAt *int32 `json:"created_at,omitempty"`

	// 批量处理对象ID
	TargetId *string `json:"target_id,omitempty"`

	// 批量处理对象基本信息
	ExtensionInfo *[]Extension `json:"extension_info,omitempty"`
}

func (o Task) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Task struct{}"
	}

	return strings.Join([]string{"Task", string(data)}, " ")
}
