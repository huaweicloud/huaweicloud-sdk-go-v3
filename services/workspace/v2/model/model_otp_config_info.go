package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OtpConfigInfo OTP辅助认证方式配置。
type OtpConfigInfo struct {

	// 认证id。
	Id *string `json:"id,omitempty"`

	// 是否启用。
	Enable *bool `json:"enable,omitempty"`

	ReceiveMode *ReceiveModeEnum `json:"receive_mode,omitempty"`

	// 辅助认证服务器地址。
	AuthUrl *string `json:"auth_url,omitempty"`

	// 认证服务接入账号。
	AppId *string `json:"app_id,omitempty"`

	// 认证服务接入密码。
	AppSecret *string `json:"app_secret,omitempty"`

	AuthServerAccessMode *AuthServerAccessMode `json:"auth_server_access_mode,omitempty"`

	// pem格式证书内容。
	CertContent *string `json:"cert_content,omitempty"`

	ApplyRule *ApplyRuleInfo `json:"apply_rule,omitempty"`

	// 要应用的用户/用户组列表。
	ApplyObjects *[]ApplyObjects `json:"apply_objects,omitempty"`
}

func (o OtpConfigInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OtpConfigInfo struct{}"
	}

	return strings.Join([]string{"OtpConfigInfo", string(data)}, " ")
}
