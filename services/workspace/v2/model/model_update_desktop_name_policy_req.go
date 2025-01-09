package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDesktopNamePolicyReq 更新桌面名称策略请求体。
type UpdateDesktopNamePolicyReq struct {

	// 策略名称，由数字、字母、中文、下划线组成，必须以字母或下划线开头，长度范围为1~30个字符。
	PolicyName *string `json:"policy_name,omitempty"`

	// 策略前缀。
	NamePrefix *string `json:"name_prefix,omitempty"`

	// 策略后缀有效位数。
	DigitNumber *int32 `json:"digit_number,omitempty"`

	// 策略后缀起始数字。
	StartNumber *int32 `json:"start_number,omitempty"`

	// 是否单用户名递增。 - 1 单用户名递增。 - 0 租户递增。
	SingleDomainUserInc *int32 `json:"single_domain_user_inc,omitempty"`

	// 是否为默认策略，true默认策略。
	IsDefaultPolicy *bool `json:"is_default_policy,omitempty"`
}

func (o UpdateDesktopNamePolicyReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDesktopNamePolicyReq struct{}"
	}

	return strings.Join([]string{"UpdateDesktopNamePolicyReq", string(data)}, " ")
}
