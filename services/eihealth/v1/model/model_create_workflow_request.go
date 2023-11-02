package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWorkflowRequest Request Object
type CreateWorkflowRequest struct {

	// 平台项目ID，您可以在平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	Body *WorkflowDto `json:"body,omitempty"`
}

func (o CreateWorkflowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkflowRequest struct{}"
	}

	return strings.Join([]string{"CreateWorkflowRequest", string(data)}, " ")
}
