package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// RadiusAuthConfig Radius认证配置信息
type RadiusAuthConfig struct {

	// Radius主机配置列表
	RadiusConfigs *[]RadiusHost `json:"radius_configs,omitempty"`

	// Radius密码
	RadiusSecret *string `json:"radius_secret,omitempty"`

	// 超时时间,取值范围1-300s
	RadiusTimeout *int32 `json:"radius_timeout,omitempty"`

	// 尝试次数,取值范围1-10
	RadiusAttemptTime *int32 `json:"radius_attempt_time,omitempty"`

	// 用户前缀中的域。0代表禁用，1代表开启域前缀，2代表开启域后缀，默认值为0。
	DomainInUserPrefix *string `json:"domain_in_user_prefix,omitempty"`

	// 认证类型。静态密码: \"STATIC\"，动态口令: \"DYNAMIC\"
	RadiusAuthType *RadiusAuthConfigRadiusAuthType `json:"radius_auth_type,omitempty"`

	// 是否启用
	Enable *bool `json:"enable,omitempty"`

	// 身份验证类型: pap或chap，选择短信验证码时才可用
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

func (o RadiusAuthConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RadiusAuthConfig struct{}"
	}

	return strings.Join([]string{"RadiusAuthConfig", string(data)}, " ")
}

type RadiusAuthConfigRadiusAuthType struct {
	value string
}

type RadiusAuthConfigRadiusAuthTypeEnum struct {
	STATIC  RadiusAuthConfigRadiusAuthType
	DYNAMIC RadiusAuthConfigRadiusAuthType
}

func GetRadiusAuthConfigRadiusAuthTypeEnum() RadiusAuthConfigRadiusAuthTypeEnum {
	return RadiusAuthConfigRadiusAuthTypeEnum{
		STATIC: RadiusAuthConfigRadiusAuthType{
			value: "STATIC",
		},
		DYNAMIC: RadiusAuthConfigRadiusAuthType{
			value: "DYNAMIC",
		},
	}
}

func (c RadiusAuthConfigRadiusAuthType) Value() string {
	return c.value
}

func (c RadiusAuthConfigRadiusAuthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RadiusAuthConfigRadiusAuthType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
