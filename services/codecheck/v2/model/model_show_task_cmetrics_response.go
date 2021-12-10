package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowTaskCmetricsResponse struct {
	// 任务id

	TaskId *string `json:"task_id,omitempty"`
	// 任务名字

	TaskName *string `json:"task_name,omitempty"`
	// 创建者id

	CreatorId *string `json:"creator_id,omitempty"`
	// 代码仓地址

	GitUrl *string `json:"git_url,omitempty"`
	// 代码仓分支

	GitBranch *string `json:"git_branch,omitempty"`
	// 上一次检查时间

	LastCheckTime *string `json:"last_check_time,omitempty"`
	// 上次执行时间

	LastExecTime *string `json:"last_exec_time,omitempty"`
	// 检查类型

	CheckType *string `json:"check_type,omitempty"`
	// 创建时间

	CreatedAt *string `json:"created_at,omitempty"`

	MetricInfo     *MetricInfo `json:"metric_info,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowTaskCmetricsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskCmetricsResponse struct{}"
	}

	return strings.Join([]string{"ShowTaskCmetricsResponse", string(data)}, " ")
}
