package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InteractionRuleGroup 互动规则库信息
type InteractionRuleGroup struct {

	// 互动规则库名称
	GroupName string `json:"group_name"`

	// 互动规则列表
	InteractionRules *[]InteractionRuleInfo `json:"interaction_rules,omitempty"`
}

func (o InteractionRuleGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InteractionRuleGroup struct{}"
	}

	return strings.Join([]string{"InteractionRuleGroup", string(data)}, " ")
}
