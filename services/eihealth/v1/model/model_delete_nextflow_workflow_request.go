package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteNextflowWorkflowRequest Request Object
type DeleteNextflowWorkflowRequest struct {

	// 平台项目ID，您可以在平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 流程id
	WorkflowId string `json:"workflow_id"`
}

func (o DeleteNextflowWorkflowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNextflowWorkflowRequest struct{}"
	}

	return strings.Join([]string{"DeleteNextflowWorkflowRequest", string(data)}, " ")
}
