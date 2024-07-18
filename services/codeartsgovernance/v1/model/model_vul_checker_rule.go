package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VulCheckerRule struct {

	// 检测项
	CheckerRule *string `json:"checker_rule,omitempty"`

	// 规范检测项条目
	CheckerEntry *string `json:"checker_entry,omitempty"`

	// 规则检测项结果
	CheckerResult *string `json:"checker_result,omitempty"`

	// 调用栈信息
	CheckerStack *string `json:"checker_stack,omitempty"`

	// 隐私声明
	PrivacyPolicyEvidences *string `json:"privacy_policy_evidences,omitempty"`

	// 是否通过
	Pass *bool `json:"pass,omitempty"`

	// 修复建议
	RectifySuggestion *string `json:"rectify_suggestion,omitempty"`
}

func (o VulCheckerRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulCheckerRule struct{}"
	}

	return strings.Join([]string{"VulCheckerRule", string(data)}, " ")
}
