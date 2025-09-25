package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityConfigWeakPwdInfo 经典弱口令的账号信息
type SecurityConfigWeakPwdInfo struct {

	// **参数解释**： 弱口令账号名称 **取值范围**： 不涉及
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**： 账号类型 **取值范围**： 不涉及
	ServiceType *string `json:"service_type,omitempty"`

	// **参数解释**： 弱口令使用时长，单位天 **取值范围**： 不涉及
	Duration *int32 `json:"duration,omitempty"`
}

func (o SecurityConfigWeakPwdInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityConfigWeakPwdInfo struct{}"
	}

	return strings.Join([]string{"SecurityConfigWeakPwdInfo", string(data)}, " ")
}
