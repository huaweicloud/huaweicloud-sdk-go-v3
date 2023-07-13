package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNatGatewaySnatRuleRequest Request Object
type UpdateNatGatewaySnatRuleRequest struct {

	// SNAT规则的ID。
	SnatRuleId string `json:"snat_rule_id"`

	Body *UpdateNatGatewaySnatRuleRequestOption `json:"body,omitempty"`
}

func (o UpdateNatGatewaySnatRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNatGatewaySnatRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateNatGatewaySnatRuleRequest", string(data)}, " ")
}
