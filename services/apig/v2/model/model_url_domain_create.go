package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UrlDomainCreate struct {

	// 最小ssl协议版本号。支持TLSv1.1或TLSv1.2
	MinSslVersion *UrlDomainCreateMinSslVersion `json:"min_ssl_version,omitempty"`

	// 是否开启http到https的重定向，false为关闭，true为开启，默认为false
	IsHttpRedirectToHttps *bool `json:"is_http_redirect_to_https,omitempty"`

	// 自定义域名。长度为0-255位的字符串，需要符合域名规范。
	UrlDomain *string `json:"url_domain,omitempty"`
}

func (o UrlDomainCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UrlDomainCreate struct{}"
	}

	return strings.Join([]string{"UrlDomainCreate", string(data)}, " ")
}

type UrlDomainCreateMinSslVersion struct {
	value string
}

type UrlDomainCreateMinSslVersionEnum struct {
	TL_SV1_1 UrlDomainCreateMinSslVersion
	TL_SV1_2 UrlDomainCreateMinSslVersion
}

func GetUrlDomainCreateMinSslVersionEnum() UrlDomainCreateMinSslVersionEnum {
	return UrlDomainCreateMinSslVersionEnum{
		TL_SV1_1: UrlDomainCreateMinSslVersion{
			value: "TLSv1.1",
		},
		TL_SV1_2: UrlDomainCreateMinSslVersion{
			value: "TLSv1.2",
		},
	}
}

func (c UrlDomainCreateMinSslVersion) Value() string {
	return c.value
}

func (c UrlDomainCreateMinSslVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UrlDomainCreateMinSslVersion) UnmarshalJSON(b []byte) error {
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
