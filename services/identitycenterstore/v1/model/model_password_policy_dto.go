package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PasswordPolicyDto struct {

	// 最小密码长度
	MinimumPasswordLength *int32 `json:"minimum_password_length,omitempty"`

	// 是否要求小写字母
	RequireLowercaseCharacters *bool `json:"require_lowercase_characters,omitempty"`

	// 是否要求数字
	RequireNumbers *bool `json:"require_numbers,omitempty"`

	// 是否要求特殊字符
	RequireSymbols *bool `json:"require_symbols,omitempty"`

	// 是否要求大写字母
	RequireUppercaseCharacters *bool `json:"require_uppercase_characters,omitempty"`

	// 密码有效期
	MaxPasswordAge *int32 `json:"max_password_age,omitempty"`

	// 密码重复使用次数，默认只支持1
	PasswordReusePrevention *int32 `json:"password_reuse_prevention,omitempty"`
}

func (o PasswordPolicyDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PasswordPolicyDto struct{}"
	}

	return strings.Join([]string{"PasswordPolicyDto", string(data)}, " ")
}
