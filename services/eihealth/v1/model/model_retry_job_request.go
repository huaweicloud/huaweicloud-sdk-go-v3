package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RetryJobRequest Request Object
type RetryJobRequest struct {

	// 医疗智能体平台项目ID，您可以在EIHealth平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// job_id
	JobId string `json:"job_id"`
}

func (o RetryJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryJobRequest struct{}"
	}

	return strings.Join([]string{"RetryJobRequest", string(data)}, " ")
}
