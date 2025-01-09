package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// LdapConfig ldap认证配置
type LdapConfig struct {

	// host
	Host *string `json:"host,omitempty"`

	// 端口,取值范围1-65535,默认389
	Port *int32 `json:"port,omitempty"`

	// base_dn
	BaseDn *string `json:"base_dn,omitempty"`

	// 管理员dn
	AdministratorDn *string `json:"administrator_dn,omitempty"`

	// 管理员密码
	AdministratorPassword *string `json:"administrator_password,omitempty"`

	// 用户dn
	UserDn *string `json:"user_dn,omitempty"`

	// 是否启用ssl
	UseSsl *bool `json:"use_ssl,omitempty"`

	// 证书
	CertContent *string `json:"cert_content,omitempty"`

	// 用户名属性
	UsernameAttribute *string `json:"username_attribute,omitempty"`

	// 用户ObjectClass
	ObjectClass *string `json:"object_class,omitempty"`

	// 安全类型
	SecurityType *LdapConfigSecurityType `json:"security_type,omitempty"`
}

func (o LdapConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LdapConfig struct{}"
	}

	return strings.Join([]string{"LdapConfig", string(data)}, " ")
}

type LdapConfigSecurityType struct {
	value string
}

type LdapConfigSecurityTypeEnum struct {
	USE_SSL LdapConfigSecurityType
	USE_TLS LdapConfigSecurityType
	CLOSE   LdapConfigSecurityType
}

func GetLdapConfigSecurityTypeEnum() LdapConfigSecurityTypeEnum {
	return LdapConfigSecurityTypeEnum{
		USE_SSL: LdapConfigSecurityType{
			value: "USE_SSL",
		},
		USE_TLS: LdapConfigSecurityType{
			value: "USE_TLS",
		},
		CLOSE: LdapConfigSecurityType{
			value: "CLOSE",
		},
	}
}

func (c LdapConfigSecurityType) Value() string {
	return c.value
}

func (c LdapConfigSecurityType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LdapConfigSecurityType) UnmarshalJSON(b []byte) error {
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
