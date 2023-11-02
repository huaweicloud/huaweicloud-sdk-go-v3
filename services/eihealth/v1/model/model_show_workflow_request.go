package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWorkflowRequest Request Object
type ShowWorkflowRequest struct {

	// 是否显示模板参数详情
	ShowParamDetail *bool `json:"Show-Param-Detail,omitempty"`

	// 平台项目ID，您可以在平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 流程id
	WorkflowId string `json:"workflow_id"`
}

func (o ShowWorkflowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkflowRequest struct{}"
	}

	return strings.Join([]string{"ShowWorkflowRequest", string(data)}, " ")
}
