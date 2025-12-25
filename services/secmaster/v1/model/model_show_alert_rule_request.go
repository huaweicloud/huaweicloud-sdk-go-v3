package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAlertRuleRequest Request Object
type ShowAlertRuleRequest struct {

	// 工作空间 ID
	WorkspaceId string `json:"workspace_id"`

	// 告警规则 ID
	RuleId string `json:"rule_id"`
}

func (o ShowAlertRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlertRuleRequest struct{}"
	}

	return strings.Join([]string{"ShowAlertRuleRequest", string(data)}, " ")
}
