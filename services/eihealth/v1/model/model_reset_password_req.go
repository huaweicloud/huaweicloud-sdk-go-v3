package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 重置密码请求体
type ResetPasswordReq struct {

	// 验证码
	Code *string `json:"code,omitempty"`

	// 认证方式
	Method *string `json:"method,omitempty"`

	// 原始密码
	OriginalPassword string `json:"original_password"`

	// 新密码
	NewPassword string `json:"new_password"`
}

func (o ResetPasswordReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetPasswordReq struct{}"
	}

	return strings.Join([]string{"ResetPasswordReq", string(data)}, " ")
}
