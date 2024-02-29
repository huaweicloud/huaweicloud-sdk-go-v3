package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchStrategyRequest Request Object
type SwitchStrategyRequest struct {

	// 策略ID
	RuleSetId string `json:"rule_set_id"`

	// 租户ID
	DomainId string `json:"domain_id"`

	Body *UpdateRuleSetStatusReq `json:"body,omitempty"`
}

func (o SwitchStrategyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchStrategyRequest struct{}"
	}

	return strings.Join([]string{"SwitchStrategyRequest", string(data)}, " ")
}
