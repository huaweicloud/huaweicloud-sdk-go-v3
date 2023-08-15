package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopNextflowJobRequest Request Object
type StopNextflowJobRequest struct {

	// 医疗智能体平台项目ID，您可以在EIHealth平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 作业id
	JobId string `json:"job_id"`
}

func (o StopNextflowJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopNextflowJobRequest struct{}"
	}

	return strings.Join([]string{"StopNextflowJobRequest", string(data)}, " ")
}
