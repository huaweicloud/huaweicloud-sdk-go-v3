package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RetryNextflowJobRequest Request Object
type RetryNextflowJobRequest struct {

	// 平台项目ID，您可以在平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 作业id
	JobId string `json:"job_id"`

	Body *RetryNextflowJobRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o RetryNextflowJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryNextflowJobRequest struct{}"
	}

	return strings.Join([]string{"RetryNextflowJobRequest", string(data)}, " ")
}
