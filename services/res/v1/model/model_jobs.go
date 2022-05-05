package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Jobs struct {

	// 类别。
	Category *string `json:"category,omitempty"`

	// 配置信息。
	ConfigInfo *string `json:"config_info,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 作业id。
	JobId string `json:"job_id"`

	// 作业名称。
	JobName *string `json:"job_name,omitempty"`

	// 作业类型。
	JobType *string `json:"job_type,omitempty"`

	// 下次调度时间。
	NextScheduleTime *int32 `json:"next_schedule_time,omitempty"`

	// 平台。
	Platform *string `json:"platform,omitempty"`

	// 资源id。
	ResourceId *string `json:"resource_id,omitempty"`

	// 调度参数。
	Schedule *string `json:"schedule,omitempty"`

	// 状态。
	Status *string `json:"status,omitempty"`

	// 工作空间id。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	JobConfig *JobConfig `json:"job_config,omitempty"`
}

func (o Jobs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Jobs struct{}"
	}

	return strings.Join([]string{"Jobs", string(data)}, " ")
}
