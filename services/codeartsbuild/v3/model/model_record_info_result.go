package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecordInfoResult 结果
type RecordInfoResult struct {

	// id
	Id *string `json:"id,omitempty"`

	// 构建工程ID,唯一对应codeci_job_id
	BuildProjectId *string `json:"build_project_id,omitempty"`

	// 构建记录ID
	BuildRecordId *string `json:"build_record_id,omitempty"`

	// 父构建记录ID
	ParentRecordId *string `json:"parent_record_id,omitempty"`

	// 项目ID
	DevcloudProjectId *string `json:"devcloud_project_id,omitempty"`

	// codeci任务ID,唯一对应build_project_id
	CodeciJobId *string `json:"codeci_job_id,omitempty"`

	// 用户ID
	UserId *string `json:"user_id,omitempty"`

	// 构建编号
	BuildNo *int32 `json:"build_no,omitempty"`

	// 每日构建编号，每日从1开始
	DailyBuildNum *string `json:"daily_build_num,omitempty"`

	// 八爪鱼任务ID
	ExecutionId *string `json:"execution_id,omitempty"`

	// 仓库名称
	RepoName *string `json:"repo_name,omitempty"`

	// 仓库id
	RepoId *string `json:"repo_id,omitempty"`

	// 仓库分支
	Branch *string `json:"branch,omitempty"`

	// 仓库tag
	Tag *string `json:"tag,omitempty"`

	// 仓库commit ID
	Commit *string `json:"commit,omitempty"`

	// 仓库commit提交信息
	CommitMessage *string `json:"commit_message,omitempty"`

	// commit创建时间
	CommitCreateTime *string `json:"commit_create_time,omitempty"`

	// 触发类型
	TriggerType *string `json:"trigger_type,omitempty"`

	// 构建类型
	BuildType *string `json:"build_type,omitempty"`

	// 构建状态
	Status *string `json:"status,omitempty"`

	// 租户ID
	DomainId *string `json:"domain_id,omitempty"`

	// 任务创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 构建下发耗时
	ScheduleTime *string `json:"schedule_time,omitempty"`

	// 构建排队耗时
	QueuedTime *string `json:"queued_time,omitempty"`

	// 开始构建时间
	StartTime *string `json:"start_time,omitempty"`

	// 八爪鱼真正开始构建时间
	RunnableTime *string `json:"runnable_time,omitempty"`

	// 构建结束时间
	FinishTime *string `json:"finish_time,omitempty"`

	// 构建时长
	Duration *float32 `json:"duration,omitempty"`

	// record状态
	RecordStatus *string `json:"record_status,omitempty"`

	// 是否使用自定义执行机
	UsePrivateSlave *int32 `json:"use_private_slave,omitempty"`

	// 租户所在region
	Region *string `json:"region,omitempty"`

	// 错误信息
	ErrMsg *string `json:"err_msg,omitempty"`

	// 构建配置类型，YAML或ACTION
	BuildConfigType *string `json:"build_config_type,omitempty"`
}

func (o RecordInfoResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecordInfoResult struct{}"
	}

	return strings.Join([]string{"RecordInfoResult", string(data)}, " ")
}
