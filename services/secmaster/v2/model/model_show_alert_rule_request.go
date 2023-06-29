package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAlertRuleRequest Request Object
type ShowAlertRuleRequest struct {

	// workspace_id
	WorkspaceId string `json:"workspace_id"`

	// rule_id
	RuleId string `json:"rule_id"`
}

func (o ShowAlertRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlertRuleRequest struct{}"
	}

	return strings.Join([]string{"ShowAlertRuleRequest", string(data)}, " ")
}
