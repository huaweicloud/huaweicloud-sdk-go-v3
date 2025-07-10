package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWorkflowTemplateRequest Request Object
type ShowWorkflowTemplateRequest struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`

	// 工作项类型
	IssueCategory string `json:"issue_category"`
}

func (o ShowWorkflowTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkflowTemplateRequest struct{}"
	}

	return strings.Join([]string{"ShowWorkflowTemplateRequest", string(data)}, " ")
}
