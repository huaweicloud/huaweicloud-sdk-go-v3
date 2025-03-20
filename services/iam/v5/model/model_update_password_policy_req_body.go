package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdatePasswordPolicyReqBody struct {

	// 同一字符连续出现的最大次数。
	MaximumConsecutiveIdenticalChars *int32 `json:"maximum_consecutive_identical_chars,omitempty"`

	// 密码最短使用时间（分钟）。
	MinimumPasswordAge *int32 `json:"minimum_password_age,omitempty"`

	// 密码最小字符数。
	MinimumPasswordLength *int32 `json:"minimum_password_length,omitempty"`

	// 密码不能与历史密码重复次数。
	PasswordReusePrevention *int32 `json:"password_reuse_prevention,omitempty"`

	// 密码是否可以是用户名或用户名的反序。
	PasswordNotUsernameOrInvert *bool `json:"password_not_username_or_invert,omitempty"`

	// 密码有效期（天）。
	PasswordValidityPeriod *int32 `json:"password_validity_period,omitempty"`

	// 至少包含字符种类的个数。
	PasswordCharCombination *int32 `json:"password_char_combination,omitempty"`

	// 是否允许IAM用户修改自己的密码，不适用于根用户。
	AllowUserToChangePassword *bool `json:"allow_user_to_change_password,omitempty"`
}

func (o UpdatePasswordPolicyReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePasswordPolicyReqBody struct{}"
	}

	return strings.Join([]string{"UpdatePasswordPolicyReqBody", string(data)}, " ")
}
