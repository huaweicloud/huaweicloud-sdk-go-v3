package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReferencedRule struct {

	// 规则ID
	RuleId *string `json:"rule_id,omitempty"`

	// 策略名称
	RuleName *string `json:"rule_name,omitempty"`

	// 引用该代答库的规则所属策略ID
	PolicyId *string `json:"policy_id,omitempty"`

	// 引用该代答库的规则所属策略名称
	PolicyName *string `json:"policy_name,omitempty"`

	// 规则检测选项类型，可选：lexicon(词库检测), compliance(合规检测), inject(注入检测), immunity_lexicon(豁免词库), supported_language(语言检测), pii(PII检测), abuse(算力滥用检测), topic_detect(话题检测)等
	DetectOptsType *string `json:"detect_opts_type,omitempty"`
}

func (o ReferencedRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReferencedRule struct{}"
	}

	return strings.Join([]string{"ReferencedRule", string(data)}, " ")
}
