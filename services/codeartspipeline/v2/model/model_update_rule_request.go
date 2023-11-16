package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRuleRequest Request Object
type UpdateRuleRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	// 规则ID
	RuleId string `json:"rule_id"`

	Body *UpdateRuleReq `json:"body,omitempty"`
}

func (o UpdateRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateRuleRequest", string(data)}, " ")
}
