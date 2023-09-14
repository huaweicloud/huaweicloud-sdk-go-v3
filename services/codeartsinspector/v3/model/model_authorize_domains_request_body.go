package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AuthorizeDomainsRequestBody struct {

	// 网站域名
	DomainName string `json:"domain_name"`

	// 认证方式:   * file - 文件认证   * auto - 一键认证   * free - 免认证，选择此项默认已阅读并了解下述使用要求           使用须知：           1、您的账号已完成实名认证，且非受限账号。           2、您确认您已获得对扫描对象进行扫描的相关合法权利。           3、您确认您的扫描行为有合法合理目的，且符合适用的法律法规要求，不得利用本服务从事任何黑灰产等非法活动。           4、若您违反上述承诺，我们有权立即终止您对本服务的使用，并要求您对我们及相关第三方因此遭受的损失进行赔偿。
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
	FREE AuthorizeDomainsRequestBodyAuthMode
}

func GetAuthorizeDomainsRequestBodyAuthModeEnum() AuthorizeDomainsRequestBodyAuthModeEnum {
	return AuthorizeDomainsRequestBodyAuthModeEnum{
		FILE: AuthorizeDomainsRequestBodyAuthMode{
			value: "file",
		},
		AUTO: AuthorizeDomainsRequestBodyAuthMode{
			value: "auto",
		},
		FREE: AuthorizeDomainsRequestBodyAuthMode{
			value: "free",
		},
	}
}

func (c AuthorizeDomainsRequestBodyAuthMode) Value() string {
	return c.value
}

func (c AuthorizeDomainsRequestBodyAuthMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AuthorizeDomainsRequestBodyAuthMode) UnmarshalJSON(b []byte) error {
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
