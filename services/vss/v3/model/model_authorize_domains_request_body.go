package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type AuthorizeDomainsRequestBody struct {
	// 域名

	DomainName string `json:"domain_name"`
	// 认证方式:   * file - 文件认证   * auto - 一键认证

	AuthMode *AuthorizeDomainsRequestBodyAuthMode `json:"auth_mode,omitempty"`
}

func (o AuthorizeDomainsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthorizeDomainsRequestBody struct{}"
	}

	return strings.Join([]string{"AuthorizeDomainsRequestBody", string(data)}, " ")
}

type AuthorizeDomainsRequestBodyAuthMode struct {
	value string
}

type AuthorizeDomainsRequestBodyAuthModeEnum struct {
	FILE AuthorizeDomainsRequestBodyAuthMode
	AUTO AuthorizeDomainsRequestBodyAuthMode
}

func GetAuthorizeDomainsRequestBodyAuthModeEnum() AuthorizeDomainsRequestBodyAuthModeEnum {
	return AuthorizeDomainsRequestBodyAuthModeEnum{
		FILE: AuthorizeDomainsRequestBodyAuthMode{
			value: "file",
		},
		AUTO: AuthorizeDomainsRequestBodyAuthMode{
			value: "auto",
		},
	}
}

func (c AuthorizeDomainsRequestBodyAuthMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AuthorizeDomainsRequestBodyAuthMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
