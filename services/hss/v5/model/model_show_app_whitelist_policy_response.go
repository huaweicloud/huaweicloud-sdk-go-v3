package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAppWhitelistPolicyResponse Response Object
type ShowAppWhitelistPolicyResponse struct {

	// 策略ID
	PolicyId *string `json:"policy_id,omitempty"`

	// 策略名称
	PolicyName *string `json:"policy_name,omitempty"`

	// **参数解释**： 进程白名单策略类型 **取值范围**: - allow：允许指定/授权进程运行 - block：阻止潜在恶意软件运行
	PolicyType *string `json:"policy_type,omitempty"`

	// **参数解释**： 是否开启阻断 **取值范围**: - true：是 - false：否
	Intercept *bool `json:"intercept,omitempty"`

	// 更新时间，毫秒
	UpdateTime *int64 `json:"update_time,omitempty"`

	// **参数解释**: 进程总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 待确认进程数 **取值范围**: 最小值0，最大值2147483647
	UnconfirmedNum *int32 `json:"unconfirmed_num,omitempty"`

	// **参数解释**: 识别可信进程数 **取值范围**: 最小值0，最大值2147483647
	TrustNum *int32 `json:"trust_num,omitempty"`

	// **参数解释**: 识别可疑进程数 **取值范围**: 最小值0，最大值2147483647
	SuspiciousNum *int32 `json:"suspicious_num,omitempty"`

	// **参数解释**: 识别恶意进程数 **取值范围**: 最小值0，最大值2147483647
	MaliciousNum *int32 `json:"malicious_num,omitempty"`

	// **参数解释**: 识别未知进程数 **取值范围**: 最小值0，最大值2147483647
	UnknownNum *int32 `json:"unknown_num,omitempty"`

	// **参数解释**： 是否自动应用策略 **取值范围**: - true：是 - false：否
	AutoApply      *bool `json:"auto_apply,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowAppWhitelistPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppWhitelistPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowAppWhitelistPolicyResponse", string(data)}, " ")
}
