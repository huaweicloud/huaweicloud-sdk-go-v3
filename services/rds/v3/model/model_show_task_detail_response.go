package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskDetailResponse Response Object
type ShowTaskDetailResponse struct {

	// 实例ID。
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例名。
	InstanceName *string `json:"instance_name,omitempty"`

	// 工作流ID。
	JobId *string `json:"job_id,omitempty"`

	// 工作流名称。
	JobName *string `json:"job_name,omitempty"`

	// 工作流创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// 工作流更新时间。
	UpdateAt *string `json:"update_at,omitempty"`

	// 工作流进度。
	Process *string `json:"process,omitempty"`

	// 工作流状态。
	Status *string `json:"status,omitempty"`

	// 子任务进度信息
	SubTaskList *[]SubTaskInfo `json:"sub_task_list,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowTaskDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowTaskDetailResponse", string(data)}, " ")
}
