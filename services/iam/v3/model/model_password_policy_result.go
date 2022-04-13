package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type PasswordPolicyResult struct {
	// 同一字符连续出现的最大次数。

	MaximumConsecutiveIdenticalChars int32 `json:"maximum_consecutive_identical_chars"`
	// 密码最大字符数。

	MaximumPasswordLength int32 `json:"maximum_password_length"`
	// 密码最短使用时间（分钟）。

	MinimumPasswordAge int32 `json:"minimum_password_age"`
	// 密码最小字符数。

	MinimumPasswordLength int32 `json:"minimum_password_length"`
	// 密码不能与历史密码重复次数。

	NumberOfRecentPasswordsDisallowed int32 `json:"number_of_recent_passwords_disallowed"`
	// 密码是否可以是用户名或用户名的反序。

	PasswordNotUsernameOrInvert bool `json:"password_not_username_or_invert"`
	// 设置密码必须包含的字符要求。

	PasswordRequirements string `json:"password_requirements"`
	// 密码有效期（天）。

	PasswordValidityPeriod int32 `json:"password_validity_period"`
	// 至少包含字符种类的个数，取值区间[2,4]。

	PasswordCharCombination int32 `json:"password_char_combination"`
}

func (o PasswordPolicyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PasswordPolicyResult struct{}"
	}

	return strings.Join([]string{"PasswordPolicyResult", string(data)}, " ")
}
