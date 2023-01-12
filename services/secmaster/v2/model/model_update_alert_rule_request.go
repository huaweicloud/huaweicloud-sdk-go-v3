package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateAlertRuleRequest struct {

	// workspace_id
	WorkspaceId string `json:"workspace_id"`

	// rule_id
	RuleId string `json:"rule_id"`

	Body *UpdateAlertRuleRequestBody `json:"body,omitempty"`
}

func (o UpdateAlertRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlertRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateAlertRuleRequest", string(data)}, " ")
}
