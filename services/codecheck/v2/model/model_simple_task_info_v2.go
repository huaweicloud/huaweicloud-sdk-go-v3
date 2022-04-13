package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SimpleTaskInfoV2 struct {
	// 任务id

	TaskId *string `json:"task_id,omitempty"`
	// 任务名字

	TaskName *string `json:"task_name,omitempty"`
	// 创建者id

	CreatorId *string `json:"creator_id,omitempty"`
	// 代码仓地址

	GitUrl *string `json:"git_url,omitempty"`
	// 代码仓分支,如果是MR模式，为源分支

	GitBranch *string `json:"git_branch,omitempty"`
	// 创建时间

	CreatedAt *string `json:"created_at,omitempty"`
	// 上一次检查时间

	LastCheckTime *string `json:"last_check_time,omitempty"`
}

func (o SimpleTaskInfoV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SimpleTaskInfoV2 struct{}"
	}

	return strings.Join([]string{"SimpleTaskInfoV2", string(data)}, " ")
}
