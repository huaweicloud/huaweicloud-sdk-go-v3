package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAlertRuleTemplateRequest Request Object
type ShowAlertRuleTemplateRequest struct {

	// 工作空间 ID。Workspace ID.
	WorkspaceId string `json:"workspace_id"`

	// 告警规则模板 ID。Alert rule template ID.
	TemplateId string `json:"template_id"`
}

func (o ShowAlertRuleTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlertRuleTemplateRequest struct{}"
	}

	return strings.Join([]string{"ShowAlertRuleTemplateRequest", string(data)}, " ")
}
