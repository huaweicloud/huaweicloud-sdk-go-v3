package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHttpGeoIpRuleRequest Request Object
type UpdateHttpGeoIpRuleRequest struct {

	// 策略id
	PolicyId string `json:"policy_id"`

	// 防护规则id
	RuleId string `json:"rule_id"`

	Body *UpdateHttpGeoIpRuleRequestBody `json:"body,omitempty"`
}

func (o UpdateHttpGeoIpRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHttpGeoIpRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateHttpGeoIpRuleRequest", string(data)}, " ")
}
