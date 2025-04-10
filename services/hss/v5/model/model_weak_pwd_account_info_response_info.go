package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WeakPwdAccountInfoResponseInfo 弱口令的账号信息
type WeakPwdAccountInfoResponseInfo struct {

	// 弱口令账号名称
	UserName *string `json:"user_name,omitempty"`

	// 账号类型，包含如下:   - system   - mysql   - redis
	ServiceType *string `json:"service_type,omitempty"`

	// 弱口令使用时长，单位天
	Duration *int32 `json:"duration,omitempty"`

	// 脱敏弱口令
	DesensitizedWeakPasswords *string `json:"desensitized_weak_passwords,omitempty"`

	// 修改建议
	Suggestion *string `json:"suggestion,omitempty"`
}

func (o WeakPwdAccountInfoResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WeakPwdAccountInfoResponseInfo struct{}"
	}

	return strings.Join([]string{"WeakPwdAccountInfoResponseInfo", string(data)}, " ")
}
