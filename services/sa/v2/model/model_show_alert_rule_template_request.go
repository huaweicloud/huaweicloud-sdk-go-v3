package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowAlertRuleTemplateRequest struct {

	// workspace_id
	WorkspaceId string `json:"workspace_id"`

	// template_id
	TemplateId string `json:"template_id"`
}

func (o ShowAlertRuleTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlertRuleTemplateRequest struct{}"
	}

	return strings.Join([]string{"ShowAlertRuleTemplateRequest", string(data)}, " ")
}
