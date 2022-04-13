package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowAomMappingRuleRequest struct {
	// 接入规则ID

	RuleId string `json:"rule_id"`
}

func (o ShowAomMappingRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAomMappingRuleRequest struct{}"
	}

	return strings.Join([]string{"ShowAomMappingRuleRequest", string(data)}, " ")
}
