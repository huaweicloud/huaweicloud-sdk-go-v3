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

	// 访问该域名绑定的http协议入方向端口，-1表示无端口且协议不支持，可使用80默认端口，其他有效端口允许的取值范围为1024~49151，需为实例已开放的HTTP协议的自定义入方向端口。  当创建域名时，该参数未填表示用默认80端口；若填写该参数，则必须同时填写https_port；若要http_port和https_port同时使用默认端口，则两个参数都不填。  当修改域名时，该参数未填表示不修改该端口。
	IngressHttpPort *int32 `json:"ingress_http_port,omitempty"`

	// 访问该域名绑定的http协议入方向端口，-1表示无端口且协议不支持，可使用443默认端口，其他有效端口允许的取值范围为1024~49151，需为实例已开放的HTTPS协议的自定义入方向端口。  当创建域名时，该参数未填表示用默认443端口；若填写该参数，则必须同时填写http_port；若要http_port和https_port同时使用默认端口，则两个参数都不填。  当修改域名时，该参数未填表示不修改该端口。
	IngressHttpsPort *int32 `json:"ingress_https_port,omitempty"`
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
