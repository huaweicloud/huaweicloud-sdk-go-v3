package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RuleItem struct {

	// 规则id
	RuleId *string `json:"rule_id,omitempty"`

	// 规则所属语言
	RuleLanguage *string `json:"rule_language,omitempty"`

	// 规则名称
	RuleName *string `json:"rule_name,omitempty"`

	// 规则问题级别
	RuleSeverity *string `json:"rule_severity,omitempty"`

	// 规则标签
	RuleTages *string `json:"rule_tages,omitempty"`

	// 规则状态0：未启用，1：已启用
	Checked *string `json:"checked,omitempty"`

	// 规则配置参数阈值相关信息
	RuleConfigList *[]RuleConfig `json:"rule_config_list,omitempty"`
}

func (o RuleItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleItem struct{}"
	}

	return strings.Join([]string{"RuleItem", string(data)}, " ")
}
