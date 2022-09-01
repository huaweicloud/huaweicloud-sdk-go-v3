package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RuleListItem struct {

	// 规则集规范分类
	RuleSet *string `json:"rule_set,omitempty" xml:"rule_set"`

	// 规则id
	RuleId *string `json:"rule_id,omitempty" xml:"rule_id"`

	// 规则所属语言
	RuleLanguage *string `json:"rule_language,omitempty" xml:"rule_language"`

	// 规则名称
	RuleName *string `json:"rule_name,omitempty" xml:"rule_name"`

	// 规则问题级别
	RuleSeverity *string `json:"rule_severity,omitempty" xml:"rule_severity"`

	// 规则标签
	RuleTages *string `json:"rule_tages,omitempty" xml:"rule_tages"`

	// 正确示例
	RightExample *string `json:"right_example,omitempty" xml:"right_example"`

	// 错误示例
	ErrorExample *string `json:"error_example,omitempty" xml:"error_example"`

	// 修改建议
	ReviseOpinion *string `json:"revise_opinion,omitempty" xml:"revise_opinion"`

	// 规则描述
	RuleDesc *string `json:"rule_desc,omitempty" xml:"rule_desc"`
}

func (o RuleListItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleListItem struct{}"
	}

	return strings.Join([]string{"RuleListItem", string(data)}, " ")
}
