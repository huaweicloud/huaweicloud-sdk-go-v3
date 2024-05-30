package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskDetailResult struct {
	InstanceInfo *InstanceInfoResult `json:"instance_info,omitempty"`

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 任务名称。
	Name *string `json:"name,omitempty"`

	// 任务状态。
	Status *string `json:"status,omitempty"`

	// 任务进度，单位：%。
	Process *string `json:"process,omitempty"`

	// 创建时间，格式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	CreatedAt *string `json:"created_at,omitempty"`

	// 结束时间，格式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	EndedAt *string `json:"ended_at,omitempty"`

	// 失败原因。
	FailReason *string `json:"fail_reason,omitempty"`
}

func (o TaskDetailResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskDetailResult struct{}"
	}

	return strings.Join([]string{"TaskDetailResult", string(data)}, " ")
}
