package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHttpGeoIpRuleRequest Request Object
type ShowHttpGeoIpRuleRequest struct {

	// 策略id
	PolicyId string `json:"policy_id"`

	// 防护规则id
	RuleId string `json:"rule_id"`
}

func (o ShowHttpGeoIpRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpGeoIpRuleRequest struct{}"
	}

	return strings.Join([]string{"ShowHttpGeoIpRuleRequest", string(data)}, " ")
}
