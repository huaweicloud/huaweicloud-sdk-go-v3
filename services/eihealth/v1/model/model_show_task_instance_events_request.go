package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskInstanceEventsRequest Request Object
type ShowTaskInstanceEventsRequest struct {

	// 平台项目ID，您可以在平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 作业id
	JobId string `json:"job_id"`

	// 子任务名称
	TaskName string `json:"task_name"`

	// 子任务的并发序号
	TaskIndex *string `json:"task_index,omitempty"`

	// 子任务实例名称
	InstanceName string `json:"instance_name"`
}

func (o ShowTaskInstanceEventsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskInstanceEventsRequest struct{}"
	}

	return strings.Join([]string{"ShowTaskInstanceEventsRequest", string(data)}, " ")
}
