package model

import (
	"encoding/json"

	"strings"
)

//
type LoginPolicyOption struct {
	// 登录提示信息，取值范围[0,240]。

	AccountValidityPeriod int32 `json:"account_validity_period"`
	// 登录提示信息。

	CustomInfoForLogin string `json:"custom_info_for_login"`
	// 帐号锁定时长（分钟），取值范围[15,30]。

	LockoutDuration int32 `json:"lockout_duration"`
	// 限定时间内登录失败次数，取值范围[3,10]。

	LoginFailedTimes int32 `json:"login_failed_times"`
	// 限定时间长度（分钟），取值范围[15,60]。

	PeriodWithLoginFailures int32 `json:"period_with_login_failures"`
	// 登录会话失效时间，取值范围[15,1440]。

	SessionTimeout int32 `json:"session_timeout"`
	// 显示最近一次的登录信息。取值范围true或false。

	ShowRecentLoginInfo bool `json:"show_recent_login_info"`
}

func (o LoginPolicyOption) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "LoginPolicyOption struct{}"
	}

	return strings.Join([]string{"LoginPolicyOption", string(data)}, " ")
}
