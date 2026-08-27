package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAuthConfigResponse Response Object
type UpdateAuthConfigResponse struct {

	// 认证配置ID。
	AuthConfigId *string `json:"auth_config_id,omitempty"`

	// 是否开启短信登录。
	SmsLoginEnabled *bool `json:"sms_login_enabled,omitempty"`

	LoginCaptcha   *LoginCaptchaConfig `json:"login_captcha,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o UpdateAuthConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAuthConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateAuthConfigResponse", string(data)}, " ")
}
