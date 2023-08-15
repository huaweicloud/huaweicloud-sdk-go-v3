package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNextflowWorkflowRequest Request Object
type ListNextflowWorkflowRequest struct {

	// 医疗智能体平台项目ID，您可以在EIHealth平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 流程名称
	Name *string `json:"name,omitempty"`
}

func (o ListNextflowWorkflowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNextflowWorkflowRequest struct{}"
	}

	return strings.Join([]string{"ListNextflowWorkflowRequest", string(data)}, " ")
}
