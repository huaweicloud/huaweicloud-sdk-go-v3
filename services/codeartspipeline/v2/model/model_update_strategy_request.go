package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateStrategyRequest Request Object
type UpdateStrategyRequest struct {

	// 策略ID
	RuleSetId string `json:"rule_set_id"`

	// 租户ID
	DomainId string `json:"domain_id"`

	Body *UpdateRuleSetReq `json:"body,omitempty"`
}

func (o UpdateStrategyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStrategyRequest struct{}"
	}

	return strings.Join([]string{"UpdateStrategyRequest", string(data)}, " ")
}
