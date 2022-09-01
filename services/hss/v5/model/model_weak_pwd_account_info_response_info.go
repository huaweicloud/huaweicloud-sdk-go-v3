package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 弱口令的账号信息
type WeakPwdAccountInfoResponseInfo struct {

	// 弱口令账号名称
	UserName *string `json:"user_name,omitempty" xml:"user_name"`

	// 账号类型
	ServiceType *string `json:"service_type,omitempty" xml:"service_type"`

	// 弱口令使用时长，单位天
	Duration *int32 `json:"duration,omitempty" xml:"duration"`
}

func (o WeakPwdAccountInfoResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WeakPwdAccountInfoResponseInfo struct{}"
	}

	return strings.Join([]string{"WeakPwdAccountInfoResponseInfo", string(data)}, " ")
}
