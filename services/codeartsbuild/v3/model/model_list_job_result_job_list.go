package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListJobResultJobList struct {

	// 任务ID
	Id *string `json:"id,omitempty"`

	// 任务名称
	JobName *string `json:"job_name,omitempty"`

	// 任务创建者
	JobCreator *string `json:"job_creator,omitempty"`

	// 用户id
	UserId *string `json:"user_id,omitempty"`

	// 用户名
	UserName *string `json:"user_name,omitempty"`

	// 用户昵称
	NickName *string `json:"nick_name,omitempty"`

	// 构建任务所在项目的ID
	ProjectId *string `json:"project_id,omitempty"`

	// 构建任务所在项目的名称
	ProjectName *string `json:"project_name,omitempty"`

	// 最新执行时间
	LastBuildTime float32 `json:"last_build_time,omitempty"`

	// 健康分值
	HealthScore *int32 `json:"health_score,omitempty"`

	// 代码来源
	SourceCode *string `json:"source_code,omitempty"`

	// 最新构建状态
	LastBuildStatus *string `json:"last_build_status,omitempty"`

	// 最新运行状态
	LastJobRunningStatus *string `json:"last_job_running_status,omitempty"`

	// 执行用户
	LastBuildUser *string `json:"last_build_user,omitempty"`

	// 执行用户ID
	LastBuildUserId *string `json:"last_build_user_id,omitempty"`

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

	// 是否有查看任务权限
	IsView *bool `json:"is_view,omitempty"`

	// 是否有执行任务权限
	IsExecute *bool `json:"is_execute,omitempty"`

	// 是否有复制任务权限
	IsCopy *bool `json:"is_copy,omitempty"`

	// 是否有禁用任务权限
	IsForbidden *bool `json:"is_forbidden,omitempty"`

	// 任务记录编号
	TaskId *string `json:"task_id,omitempty"`

	// 代码分支
	CodeBranch *string `json:"code_branch,omitempty"`

	// 代码提交ID
	CommitId *string `json:"commit_id,omitempty"`

	// 触发类型
	TriggerType *string `json:"trigger_type,omitempty"`

	// 执行时间
	BuildTime float32 `json:"build_time,omitempty"`

	// 代码源地址
	ScmWebUrl *string `json:"scm_web_url,omitempty"`

	// 仓库类别，Repo、Github等
	ScmType *string `json:"scm_type,omitempty"`

	// repo的id
	RepoId *string `json:"repo_id,omitempty"`

	// 代码提交记录信息地址（代码源为Repo)
	CommitDetailUrl *string `json:"commit_detail_url,omitempty"`

	// 构建编号
	BuildNumber *string `json:"build_number,omitempty"`

	// 禁用消息
	ForbiddenMsg *string `json:"forbidden_msg,omitempty"`

	// 构建工程ID,唯一对应codeci_job_id
	BuildProjectId *string `json:"build_project_id,omitempty"`

	// 构建类别
	BuildType *string `json:"build_type,omitempty"`

	// 仓库tag
	Tag *string `json:"tag,omitempty"`

	// 使用项目权限
	ProjectPermissionSwitch *bool `json:"project_permission_switch,omitempty"`
}

func (o ListJobResultJobList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobResultJobList struct{}"
	}

	return strings.Join([]string{"ListJobResultJobList", string(data)}, " ")
}
