package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 增量检查相关的参数
type IncConfigV2 struct {
	// 需要关联的父任务ID，流水线创建或MR创建任务需要该参数

	ParentTaskId *string `json:"parent_task_id,omitempty"`
	// 增量检查代码源分支

	GitSourceBranch *string `json:"git_source_branch,omitempty"`
	// 增量检查代码目标分支

	GitTargetBranch *string `json:"git_target_branch,omitempty"`
	// MR唯一标示ID

	MergeId *string `json:"merge_id,omitempty"`
	// webhook触发事件类型,merge_request/push_request

	EventType *string `json:"event_type,omitempty"`
	// webhook事件状态，open/close/update

	Action *string `json:"action,omitempty"`
	// MR标题

	Title *string `json:"title,omitempty"`
}

func (o IncConfigV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IncConfigV2 struct{}"
	}

	return strings.Join([]string{"IncConfigV2", string(data)}, " ")
}
