package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateAutoJobRequest struct {

	// 医疗智能体平台项目ID，您可以在EIHealth平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 自动作业id
	AutoJobId string `json:"auto_job_id"`

	Body *CreateAutJobReq `json:"body,omitempty"`
}

func (o UpdateAutoJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAutoJobRequest struct{}"
	}

	return strings.Join([]string{"UpdateAutoJobRequest", string(data)}, " ")
}
