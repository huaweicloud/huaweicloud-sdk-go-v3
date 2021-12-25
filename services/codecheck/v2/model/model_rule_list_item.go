package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RuleListItem struct {
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
	// 正确示例

	RightExample *string `json:"right_example,omitempty"`
	// 错误示例

	ErrorExample *string `json:"error_example,omitempty"`
	// 修改建议

	ReviseOpinion *string `json:"revise_opinion,omitempty"`
	// 规则描述

	RuleDesc *string `json:"rule_desc,omitempty"`
}

func (o RuleListItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleListItem struct{}"
	}

	return strings.Join([]string{"RuleListItem", string(data)}, " ")
}
