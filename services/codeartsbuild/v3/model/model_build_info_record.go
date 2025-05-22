package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BuildInfoRecord 构建历史详情
type BuildInfoRecord struct {

	// 构建编号
	Number *int32 `json:"number,omitempty"`

	// 执行时间
	BuildTime *int32 `json:"build_time,omitempty"`

	// 开始时间，时间戳
	StartTime *int64 `json:"start_time,omitempty"`

	// 运行状态
	JobRunningStatus *string `json:"job_running_status,omitempty"`

	// 任务状态
	State *string `json:"state,omitempty"`

	// IAM用户ID
	UserId *string `json:"user_id,omitempty"`

	// 触发构建用户
	Executor *string `json:"executor,omitempty"`

	// 用户名称
	Nickname *string `json:"nickname,omitempty"`

	// 构建编号，每日从1开始
	DailyBuildNumber *string `json:"daily_build_number,omitempty"`

	// 触发类型
	TriggerType *string `json:"trigger_type,omitempty"`

	// 执行时间
	CostTime *int32 `json:"cost_time,omitempty"`

	// 代码提交的commit id
	CommitId *string `json:"commit_id,omitempty"`

	CommitInfo *BuildInfoRecordCommitInfo `json:"commit_info,omitempty"`

	// 构建类型
	BuildType *string `json:"build_type,omitempty"`

	// 代码仓分支
	CodeBranch *string `json:"code_branch,omitempty"`

	// 代码源类型
	ScmType *string `json:"scm_type,omitempty"`

	// 代码源地址
	ScmWebUrl *string `json:"scm_web_url,omitempty"`

	// 代码提交记录信息地址（代码源为Repo)
	CommitDetailUrl *string `json:"commit_detail_url,omitempty"`
}

func (o BuildInfoRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BuildInfoRecord struct{}"
	}

	return strings.Join([]string{"BuildInfoRecord", string(data)}, " ")
}
