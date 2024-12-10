package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePasswordRequestBody 生成随机密码入参
type CreatePasswordRequestBody struct {

	// 策略名称
	PasswordPolicyName *string `json:"password_policy_name,omitempty"`

	// 长度，默认32
	Length *int32 `json:"length,omitempty"`

	// 需要排除的字符
	ExcludeCharacters *string `json:"exclude_characters,omitempty"`

	// 排除小写字母，默认false
	ExcludeLowercase *bool `json:"exclude_lowercase,omitempty"`

	// 排除大写字母，默认false
	ExcludeUppercase *bool `json:"exclude_uppercase,omitempty"`

	// 排除数字，默认false
	ExcludeNumbers *bool `json:"exclude_numbers,omitempty"`

	// 排除符号，默认false
	ExcludePunctuation *bool `json:"exclude_punctuation,omitempty"`

	// 包含空格，默认false
	IncludeSpace *bool `json:"include_space,omitempty"`

	// 需要每个包含的类型，默认false
	RequireEachIncludedType *bool `json:"require_each_included_type,omitempty"`
}

func (o CreatePasswordRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePasswordRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePasswordRequestBody", string(data)}, " ")
}
