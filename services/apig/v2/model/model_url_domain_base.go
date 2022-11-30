package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UrlDomainBase struct {

	// 最小ssl协议版本号。支持TLSv1.1或TLSv1.2
	MinSslVersion *UrlDomainBaseMinSslVersion `json:"min_ssl_version,omitempty"`

	// 是否开启http到https的重定向，false为关闭，true为开启，默认为false
	IsHttpRedirectToHttps *bool `json:"is_http_redirect_to_https,omitempty"`
}

func (o UrlDomainBase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UrlDomainBase struct{}"
	}

	return strings.Join([]string{"UrlDomainBase", string(data)}, " ")
}

type UrlDomainBaseMinSslVersion struct {
	value string
}

type UrlDomainBaseMinSslVersionEnum struct {
	TL_SV1_1 UrlDomainBaseMinSslVersion
	TL_SV1_2 UrlDomainBaseMinSslVersion
}

func GetUrlDomainBaseMinSslVersionEnum() UrlDomainBaseMinSslVersionEnum {
	return UrlDomainBaseMinSslVersionEnum{
		TL_SV1_1: UrlDomainBaseMinSslVersion{
			value: "TLSv1.1",
		},
		TL_SV1_2: UrlDomainBaseMinSslVersion{
			value: "TLSv1.2",
		},
	}
}

func (c UrlDomainBaseMinSslVersion) Value() string {
	return c.value
}

func (c UrlDomainBaseMinSslVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UrlDomainBaseMinSslVersion) UnmarshalJSON(b []byte) error {
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
