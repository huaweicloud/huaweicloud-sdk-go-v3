package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateLoginPolicyReqBody struct {

	// 如果IAM用户在该值设置的有效期（天）内未登录，则被停用，不适用于根用户。
	UserValidityPeriod *int32 `json:"user_validity_period,omitempty"`

	// 登录提示信息，不能包含特定字符\"@\"、\"#\"、\"%\"、\"&\"、\"<\"、\">\"、\"\\\\\"、\"$\"、\"^\"和\"*\"的字符串。
	CustomInfoForLogin *string `json:"custom_info_for_login,omitempty"`

	// IAM用户登录锁定时长（分钟）。
	LockoutDuration *int32 `json:"lockout_duration,omitempty"`

	// 限定时间内登录失败次数。
	LoginFailedTimes *int32 `json:"login_failed_times,omitempty"`

	// 限定时间长度（分钟）。
	PeriodWithLoginFailures *int32 `json:"period_with_login_failures,omitempty"`

	// 登录会话失效时间。
	SessionTimeout *int32 `json:"session_timeout,omitempty"`

	// 是否显示最近一次的登录信息。
	ShowRecentLoginInfo *bool `json:"show_recent_login_info,omitempty"`

	// 允许访问的IP地址或网段。
	AllowAddressNetmasks *[]AllowAddressNetmask `json:"allow_address_netmasks,omitempty"`

	// 允许访问的IP地址区间。
	AllowIpRanges *[]AllowIpRange `json:"allow_ip_ranges,omitempty"`
}

func (o UpdateLoginPolicyReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLoginPolicyReqBody struct{}"
	}

	return strings.Join([]string{"UpdateLoginPolicyReqBody", string(data)}, " ")
}
