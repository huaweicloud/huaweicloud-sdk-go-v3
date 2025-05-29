package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BuildRecord 项目列表
type BuildRecord struct {

	// 唯一标识
	Id *string `json:"id,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 状态码
	StatusCode *int32 `json:"status_code,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 等待时间
	ScheduleTime *string `json:"schedule_time,omitempty"`

	// 排队时间
	QueuedTime *string `json:"queued_time,omitempty"`

	// 开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 完成时间
	FinishTime *string `json:"finish_time,omitempty"`

	// 持续时间
	Duration *int32 `json:"duration,omitempty"`

	// 构建时间
	BuildDuration *int32 `json:"build_duration,omitempty"`

	// 等待时间
	PendingDuration *int32 `json:"pending_duration,omitempty"`

	// 工程ID
	ProjectId *string `json:"project_id,omitempty"`

	// 子任务名称
	DisplayName *string `json:"display_name,omitempty"`

	// 触发者名称
	TriggerName *string `json:"trigger_name,omitempty"`

	// 分组名
	GroupName *string `json:"group_name,omitempty"`

	// 八爪鱼任务ID
	ExecutionId *string `json:"execution_id,omitempty"`

	// 构建执行参数列表
	Parameters *[]BuildRecordParameters `json:"parameters,omitempty"`

	// 仓库地址
	Repository *string `json:"repository,omitempty"`

	// 分支名
	Branch *string `json:"branch,omitempty"`

	// commitId
	Revision *string `json:"revision,omitempty"`

	// yaml路径
	BuildYmlPath *string `json:"build_yml_path,omitempty"`

	// yaml地址
	BuildYmlUrl *string `json:"build_yml_url,omitempty"`

	// 构建每日编号
	DailyBuildNumber *string `json:"daily_build_number,omitempty"`

	BuildRecordType *BuildRecordBuildRecordType `json:"build_record_type,omitempty"`

	// 触发类型
	TriggerType *string `json:"trigger_type,omitempty"`

	// 代码源类型
	ScmType *string `json:"scm_type,omitempty"`

	// 代码源地址
	ScmWebUrl *string `json:"scm_web_url,omitempty"`

	// 用户id
	UserId *string `json:"user_id,omitempty"`

	// 构建编码
	BuildNo *string `json:"build_no,omitempty"`

	// 构建每日编号
	DailyBuildNo *string `json:"daily_build_no,omitempty"`

	// 构建类型
	DevCloudBuildType *string `json:"dev_cloud_build_type,omitempty"`
}

func (o BuildRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BuildRecord struct{}"
	}

	return strings.Join([]string{"BuildRecord", string(data)}, " ")
}
