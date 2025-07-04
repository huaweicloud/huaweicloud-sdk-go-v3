package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Job struct {

	// 任务ID
	Id *string `json:"id,omitempty"`

	// 任务名称
	JobName *string `json:"job_name,omitempty"`

	// 任务创建者
	JobCreator *string `json:"job_creator,omitempty"`

	// 用户名称
	UserName *string `json:"user_name,omitempty"`

	// 最新执行时间
	LastBuildTime float32 `json:"last_build_time,omitempty"`

	// 健康分值
	HealthScore *int32 `json:"health_score,omitempty"`

	// 代码来源
	SourceCode *string `json:"source_code,omitempty"`

	// 最新构建状态
	LastBuildStatus *string `json:"last_build_status,omitempty"`

	// 是否已结束
	IsFinished *bool `json:"is_finished,omitempty"`

	// 是否已禁用
	Disabled *bool `json:"disabled,omitempty"`

	// 是否已收藏
	Favorite *bool `json:"favorite,omitempty"`

	// 是否有修改任务权限
	IsModify *bool `json:"is_modify,omitempty"`

	// 是否有删除任务权限
	IsDelete *bool `json:"is_delete,omitempty"`

	// 是否有执行任务权限
	IsExecute *bool `json:"is_execute,omitempty"`

	// 是否有复制任务权限
	IsCopy *bool `json:"is_copy,omitempty"`

	// 是否有禁用任务权限
	IsForbidden *bool `json:"is_forbidden,omitempty"`

	// 是否有查看任务权限
	IsView *bool `json:"is_view,omitempty"`

	// 最后一次构建用户
	LastBuildUser *string `json:"last_build_user,omitempty"`

	// 触发类型
	TriggerType *string `json:"trigger_type,omitempty"`

	// 构建时间
	BuildTime *int64 `json:"build_time,omitempty"`

	// 代码仓web地址
	ScmWebUrl *string `json:"scm_web_url,omitempty"`

	// 代码仓类型
	ScmType *string `json:"scm_type,omitempty"`

	// 代码仓ID
	RepoId *string `json:"repo_id,omitempty"`

	// 构建项目ID
	BuildProjectId *string `json:"build_project_id,omitempty"`

	// 最后一次构建时间
	LastJobRunningStatus *string `json:"last_job_running_status,omitempty"`

	// 最后一次构建用户ID
	LastBuildUserId *string `json:"last_build_user_id,omitempty"`
}

func (o Job) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Job struct{}"
	}

	return strings.Join([]string{"Job", string(data)}, " ")
}
