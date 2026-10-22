package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBotMRuleDefenseStrategyRequest Request Object
type UpdateBotMRuleDefenseStrategyRequest struct {

	// policyId
	PolicyId string `json:"policy_id"`

	Body *UpdateBotMRuleDefenseStrategyRequestBody `json:"body,omitempty"`
}

func (o UpdateBotMRuleDefenseStrategyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBotMRuleDefenseStrategyRequest struct{}"
	}

	return strings.Join([]string{"UpdateBotMRuleDefenseStrategyRequest", string(data)}, " ")
}
