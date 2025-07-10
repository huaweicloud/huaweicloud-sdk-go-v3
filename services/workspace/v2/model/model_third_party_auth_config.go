package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ThirdPartyAuthConfig 第三方认证配置信息。
type ThirdPartyAuthConfig struct {

	// 更新认证配置类型，认证类型为第三方单点登录时使用。
	UpdateType *string `json:"update_type,omitempty"`

	// 是否启用。
	Enable *bool `json:"enable,omitempty"`

	// 更新认证配置对象，认证类型为第三方单点登录时使用。
	UpdateObject *string `json:"update_object,omitempty"`

	// 认证类型。
	AuthType *string `json:"auth_type,omitempty"`

	ClientInterfaceConfig *InterfacesConfig `json:"client_interface_config,omitempty"`

	ServerInterfaceConfig *InterfacesConfig `json:"server_interface_config,omitempty"`

	// 更新认证配置类型，认证类型为第三方密码时使用。ADD代表新增，UPDATE代表修改，DELETE代表删除。
	ThirdPasswordUpdateType *ThirdPartyAuthConfigThirdPasswordUpdateType `json:"third_password_update_type,omitempty"`

	// 自定义接口配置。
	CustomDefinition *string `json:"custom_definition,omitempty"`

	// oauth2配置。
	OauthConfigs *string `json:"oauth_configs,omitempty"`

	// 单点登录配置信息列表。
	LdapConfigs *[]LdapConfig `json:"ldap_configs,omitempty"`

	// 更新认证配置对象，认证类型为第三方密码时使用。
	ThirdPasswordName *string `json:"third_password_name,omitempty"`
}

func (o ThirdPartyAuthConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ThirdPartyAuthConfig struct{}"
	}

	return strings.Join([]string{"ThirdPartyAuthConfig", string(data)}, " ")
}

type ThirdPartyAuthConfigThirdPasswordUpdateType struct {
	value string
}

type ThirdPartyAuthConfigThirdPasswordUpdateTypeEnum struct {
	ADD    ThirdPartyAuthConfigThirdPasswordUpdateType
	UPDATE ThirdPartyAuthConfigThirdPasswordUpdateType
	DELETE ThirdPartyAuthConfigThirdPasswordUpdateType
}

func GetThirdPartyAuthConfigThirdPasswordUpdateTypeEnum() ThirdPartyAuthConfigThirdPasswordUpdateTypeEnum {
	return ThirdPartyAuthConfigThirdPasswordUpdateTypeEnum{
		ADD: ThirdPartyAuthConfigThirdPasswordUpdateType{
			value: "ADD",
		},
		UPDATE: ThirdPartyAuthConfigThirdPasswordUpdateType{
			value: "UPDATE",
		},
		DELETE: ThirdPartyAuthConfigThirdPasswordUpdateType{
			value: "DELETE",
		},
	}
}

func (c ThirdPartyAuthConfigThirdPasswordUpdateType) Value() string {
	return c.value
}

func (c ThirdPartyAuthConfigThirdPasswordUpdateType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ThirdPartyAuthConfigThirdPasswordUpdateType) UnmarshalJSON(b []byte) error {
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
