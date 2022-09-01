package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowTaskCmetricsResponse struct {

	// 任务id
	TaskId *string `json:"task_id,omitempty" xml:"task_id"`

	// 任务名字
	TaskName *string `json:"task_name,omitempty" xml:"task_name"`

	// 创建者id
	CreatorId *string `json:"creator_id,omitempty" xml:"creator_id"`

	// 代码仓地址
	GitUrl *string `json:"git_url,omitempty" xml:"git_url"`

	// 代码仓分支
	GitBranch *string `json:"git_branch,omitempty" xml:"git_branch"`

	// 上一次检查时间
	LastCheckTime *string `json:"last_check_time,omitempty" xml:"last_check_time"`

	// 上次执行时间
	LastExecTime *string `json:"last_exec_time,omitempty" xml:"last_exec_time"`

	// 检查类型
	CheckType *string `json:"check_type,omitempty" xml:"check_type"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at"`

	MetricInfo     *MetricInfo `json:"metric_info,omitempty" xml:"metric_info"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowTaskCmetricsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskCmetricsResponse struct{}"
	}

	return strings.Join([]string{"ShowTaskCmetricsResponse", string(data)}, " ")
}
