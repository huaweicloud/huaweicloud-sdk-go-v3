package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LoginCaptchaConfig 滑块验证码相关配置。
type LoginCaptchaConfig struct {

	// 是否开启滑块验证码。
	Enabled *bool `json:"enabled,omitempty"`

	// 用户登录失败 trigger_threshold 次后开始要求验证码认证。验证码不启用时无意义，启用时若不传默认为 3。
	TriggerThreshold *int32 `json:"trigger_threshold,omitempty"`
}

func (o LoginCaptchaConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoginCaptchaConfig struct{}"
	}

	return strings.Join([]string{"LoginCaptchaConfig", string(data)}, " ")
}
