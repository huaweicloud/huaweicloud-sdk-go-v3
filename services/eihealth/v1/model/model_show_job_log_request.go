package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobLogRequest Request Object
type ShowJobLogRequest struct {

	// 平台项目ID，您可以在平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 作业id
	JobId string `json:"job_id"`

	// 子任务名称
	TaskName string `json:"task_name"`

	// 子任务并发的序号
	TaskIndex *string `json:"task_index,omitempty"`
}

func (o ShowJobLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobLogRequest struct{}"
	}

	return strings.Join([]string{"ShowJobLogRequest", string(data)}, " ")
}
