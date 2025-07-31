package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WeakPwdAccountInfoResponseInfo 弱口令的账号信息
type WeakPwdAccountInfoResponseInfo struct {

	// **参数解释**: 弱口令账号名称 **取值范围**: 不涉及
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**: 账号类型 **取值范围**: - system : 系统账号 - mysql  : mysql账号 - redis  : redis账号
	ServiceType *string `json:"service_type,omitempty"`

	// **参数解释**: 弱口令使用时长，单位天 **取值范围**: 不涉及
	Duration *int32 `json:"duration,omitempty"`

	// **参数解释**: 脱敏弱口令 **取值范围**: 不涉及
	DesensitizedWeakPasswords *string `json:"desensitized_weak_passwords,omitempty"`

	// **参数解释**: 修改建议 **取值范围**: 不涉及
	Suggestion *string `json:"suggestion,omitempty"`
}

func (o WeakPwdAccountInfoResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WeakPwdAccountInfoResponseInfo struct{}"
	}

	return strings.Join([]string{"WeakPwdAccountInfoResponseInfo", string(data)}, " ")
}
