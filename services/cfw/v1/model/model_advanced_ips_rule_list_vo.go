package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AdvancedIpsRuleListVo struct {
	AdvancedIpsRules *[]AdvancedIpsRuleVo `json:"advanced_ips_rules,omitempty"`

	Total *int32 `json:"total,omitempty"`
}

func (o AdvancedIpsRuleListVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AdvancedIpsRuleListVo struct{}"
	}

	return strings.Join([]string{"AdvancedIpsRuleListVo", string(data)}, " ")
}
