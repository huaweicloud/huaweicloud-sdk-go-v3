package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AdvancedIpsRuleListVo struct {

	// 频率ips规则列表
	AdvancedIpsRules *[]AdvancedIpsRuleVo `json:"advanced_ips_rules,omitempty"`

	// 频率ips规则总数
	Total *int32 `json:"total,omitempty"`
}

func (o AdvancedIpsRuleListVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AdvancedIpsRuleListVo struct{}"
	}

	return strings.Join([]string{"AdvancedIpsRuleListVo", string(data)}, " ")
}
