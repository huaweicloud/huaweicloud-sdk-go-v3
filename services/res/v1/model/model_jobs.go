package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Jobs struct {

	// 类别。
	Category *string `json:"category,omitempty" xml:"category"`

	// 配置信息。
	ConfigInfo *string `json:"config_info,omitempty" xml:"config_info"`

	// 描述。
	Description *string `json:"description,omitempty" xml:"description"`

	// 作业id。
	JobId string `json:"job_id" xml:"job_id"`

	// 作业名称。
	JobName *string `json:"job_name,omitempty" xml:"job_name"`

	// 作业类型。
	JobType *string `json:"job_type,omitempty" xml:"job_type"`

	// 下次调度时间。
	NextScheduleTime *int32 `json:"next_schedule_time,omitempty" xml:"next_schedule_time"`

	// 平台。
	Platform *string `json:"platform,omitempty" xml:"platform"`

	// 资源id。
	ResourceId *string `json:"resource_id,omitempty" xml:"resource_id"`

	// 调度参数。
	Schedule *string `json:"schedule,omitempty" xml:"schedule"`

	// 状态。
	Status *string `json:"status,omitempty" xml:"status"`

	// 工作空间id。
	WorkspaceId *string `json:"workspace_id,omitempty" xml:"workspace_id"`

	JobConfig *JobConfig `json:"job_config,omitempty" xml:"job_config"`
}

func (o Jobs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Jobs struct{}"
	}

	return strings.Join([]string{"Jobs", string(data)}, " ")
}
