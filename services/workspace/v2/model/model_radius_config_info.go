package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RadiusConfigInfo Radius配置信息
type RadiusConfigInfo struct {

	// Radius主机信息
	RadiusConfigs *[]RadiusHost `json:"radius_configs,omitempty"`

	// Radius密码
	RadiusSecret *string `json:"radius_secret,omitempty"`

	// 超时
	RadiusTimeout *int32 `json:"radius_timeout,omitempty"`

	// 尝试次数
	RadiusAttemptTime *int32 `json:"radius_attempt_time,omitempty"`

	// 用户前缀中的域，0代表禁用，1代表开启域前缀，2代表开启域后缀，默认值为0。
	DomainInUserPrefix *string `json:"domain_in_user_prefix,omitempty"`

	// 认证类型，STATIC代表静态密码，DYNAMIC代表动态口令。
	RadiusAuthType *string `json:"radius_auth_type,omitempty"`

	// 是否启用
	Enable *bool `json:"enable,omitempty"`

	// 认证协议，pap或chap，选择短信验证码时才可用。
	AuthProtocol *string `json:"auth_protocol,omitempty"`

	// 验证码开关，默认开启
	VerificationCodeDisplayEnable *bool `json:"verification_code_display_enable,omitempty"`

	// 验证码失败次数
	VerificationCodeConditions *int32 `json:"verification_code_conditions,omitempty"`

	// 过滤短信拓展字段
	ObtainCodeAttributes *interface{} `json:"obtain_code_attributes,omitempty"`

	// 过滤短信拓展字段
	CheckCodeAttributes *interface{} `json:"check_code_attributes,omitempty"`
}

func (o RadiusConfigInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RadiusConfigInfo struct{}"
	}

	return strings.Join([]string{"RadiusConfigInfo", string(data)}, " ")
}
