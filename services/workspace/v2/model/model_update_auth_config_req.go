package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAuthConfigReq 增量更新认证配置请求。
type UpdateAuthConfigReq struct {

	// 是否开启短信登录。
	SmsLoginEnabled *bool `json:"sms_login_enabled,omitempty"`

	LoginCaptcha *LoginCaptchaConfig `json:"login_captcha,omitempty"`
}

func (o UpdateAuthConfigReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAuthConfigReq struct{}"
	}

	return strings.Join([]string{"UpdateAuthConfigReq", string(data)}, " ")
}
