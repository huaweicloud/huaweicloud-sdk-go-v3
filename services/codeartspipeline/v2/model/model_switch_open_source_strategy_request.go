package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchOpenSourceStrategyRequest Request Object
type SwitchOpenSourceStrategyRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	// 策略ID
	RuleSetId string `json:"rule_set_id"`

	Body *UpdateRuleSetStatusReq `json:"body,omitempty"`
}

func (o SwitchOpenSourceStrategyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchOpenSourceStrategyRequest struct{}"
	}

	return strings.Join([]string{"SwitchOpenSourceStrategyRequest", string(data)}, " ")
}
