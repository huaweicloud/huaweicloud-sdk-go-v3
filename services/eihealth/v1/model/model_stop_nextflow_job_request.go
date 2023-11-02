package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopNextflowJobRequest Request Object
type StopNextflowJobRequest struct {

	// 平台项目ID，您可以在平台单击所需的项目名称，进入项目设置页面查看。
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
