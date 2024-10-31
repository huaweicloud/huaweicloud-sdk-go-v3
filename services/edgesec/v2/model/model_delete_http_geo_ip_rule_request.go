package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHttpGeoIpRuleRequest Request Object
type DeleteHttpGeoIpRuleRequest struct {

	// 策略id
	PolicyId string `json:"policy_id"`

	// 防护规则id
	RuleId string `json:"rule_id"`
}

func (o DeleteHttpGeoIpRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHttpGeoIpRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteHttpGeoIpRuleRequest", string(data)}, " ")
}
