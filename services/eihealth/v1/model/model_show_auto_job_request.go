package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutoJobRequest Request Object
type ShowAutoJobRequest struct {

	// 医疗智能体平台项目ID，您可以在EIHealth平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 自动作业id
	AutoJobId string `json:"auto_job_id"`
}

func (o ShowAutoJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoJobRequest struct{}"
	}

	return strings.Join([]string{"ShowAutoJobRequest", string(data)}, " ")
}
