package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAppRuleRequest Request Object
type UpdateAppRuleRequest struct {

	// 规则ID。
	RuleId string `json:"rule_id"`

	Body *UpdateAppRuleReq `json:"body,omitempty"`
}

func (o UpdateAppRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAppRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateAppRuleRequest", string(data)}, " ")
}
