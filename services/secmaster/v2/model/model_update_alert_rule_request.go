package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAlertRuleRequest Request Object
type UpdateAlertRuleRequest struct {

	// 工作空间 ID。Workspace ID.
	WorkspaceId string `json:"workspace_id"`

	// 告警规则 ID。Alert rule ID.
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
