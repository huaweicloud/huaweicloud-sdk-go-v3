package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DesktopNamePolicyInfo 桌面名称策略信息。
type DesktopNamePolicyInfo struct {

	// 策略id。
	PolicyId *string `json:"policy_id,omitempty"`

	// 策略名称。
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

	// 是否包含用户名的桌面名称策略，true包含。
	IsContainUser *bool `json:"is_contain_user,omitempty"`
}

func (o DesktopNamePolicyInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DesktopNamePolicyInfo struct{}"
	}

	return strings.Join([]string{"DesktopNamePolicyInfo", string(data)}, " ")
}
