package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityConfigPwdPolicyInfo 口令复杂度策略
type SecurityConfigPwdPolicyInfo struct {

	// **参数解释**： 口令最小长度策略是否满足要求 **取值范围**： - true：满足要求 - false：不满足要求
	MinLength *bool `json:"min_length,omitempty"`

	// **参数解释**： 大写字母策略是否满足要求 **取值范围**： - true：满足要求 - false：不满足要求
	UppercaseLetter *bool `json:"uppercase_letter,omitempty"`

	// **参数解释**： 小写字母策略是否满足要求 **取值范围**： - true：满足要求 - false：不满足要求
	LowercaseLetter *bool `json:"lowercase_letter,omitempty"`

	// **参数解释**： 数字策略是否满足要求 **取值范围**： - true：满足要求 - false：不满足要求不涉及
	Number *bool `json:"number,omitempty"`

	// **参数解释**： 特殊字符策略是否满足要求 **取值范围**： - true：满足要求 - false：不满足要求
	SpecialCharacter *bool `json:"special_character,omitempty"`

	// **参数解释**： 修改建议 **取值范围**： 不涉及
	Suggestion *string `json:"suggestion,omitempty"`
}

func (o SecurityConfigPwdPolicyInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityConfigPwdPolicyInfo struct{}"
	}

	return strings.Join([]string{"SecurityConfigPwdPolicyInfo", string(data)}, " ")
}
