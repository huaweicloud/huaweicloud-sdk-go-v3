package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UrlDomain struct {

	// 域名编号
	Id *string `json:"id,omitempty"`

	// 访问域名
	Domain *string `json:"domain,omitempty"`

	// 域名cname状态： - 1：未解析 - 2：解析中 - 3：解析成功 - 4：解析失败
	CnameStatus *int32 `json:"cname_status,omitempty"`

	// SSL证书编号
	SslId *string `json:"ssl_id,omitempty"`

	// SSL证书名称
	SslName *string `json:"ssl_name,omitempty"`

	// 最小ssl协议版本号。支持TLSv1.1或TLSv1.2
	MinSslVersion *UrlDomainMinSslVersion `json:"min_ssl_version,omitempty"`

	// 是否开启客户端证书校验。只有绑定证书时，该参数才生效。当绑定证书存在trusted_root_ca时，默认开启；当绑定证书不存在trusted_root_ca时，默认关闭。
	VerifiedClientCertificateEnabled *bool `json:"verified_client_certificate_enabled,omitempty"`

	// 是否存在信任的根证书CA。当绑定证书存在trusted_root_ca时为true。
	IsHasTrustedRootCa *bool `json:"is_has_trusted_root_ca,omitempty"`

	// 访问该域名绑定的http协议入方向端口，-1表示无端口且协议不支持，可使用80默认端口，其他有效端口允许的取值范围为1024~49151，需为实例已开放的HTTP协议的自定义入方向端口。  当创建域名时，该参数未填表示用默认80端口；若填写该参数，则必须同时填写https_port；若要http_port和https_port同时使用默认端口，则两个参数都不填。  当修改域名时，该参数未填表示不修改该端口。
	IngressHttpPort *int32 `json:"ingress_http_port,omitempty"`

	// 访问该域名绑定的http协议入方向端口，-1表示无端口且协议不支持，可使用443默认端口，其他有效端口允许的取值范围为1024~49151，需为实例已开放的HTTPS协议的自定义入方向端口。  当创建域名时，该参数未填表示用默认443端口；若填写该参数，则必须同时填写http_port；若要http_port和https_port同时使用默认端口，则两个参数都不填。  当修改域名时，该参数未填表示不修改该端口。
	IngressHttpsPort *int32 `json:"ingress_https_port,omitempty"`
}

func (o UrlDomain) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UrlDomain struct{}"
	}

	return strings.Join([]string{"UrlDomain", string(data)}, " ")
}

type UrlDomainMinSslVersion struct {
	value string
}

type UrlDomainMinSslVersionEnum struct {
	TL_SV1_1 UrlDomainMinSslVersion
	TL_SV1_2 UrlDomainMinSslVersion
}

func GetUrlDomainMinSslVersionEnum() UrlDomainMinSslVersionEnum {
	return UrlDomainMinSslVersionEnum{
		TL_SV1_1: UrlDomainMinSslVersion{
			value: "TLSv1.1",
		},
		TL_SV1_2: UrlDomainMinSslVersion{
			value: "TLSv1.2",
		},
	}
}

func (c UrlDomainMinSslVersion) Value() string {
	return c.value
}

func (c UrlDomainMinSslVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UrlDomainMinSslVersion) UnmarshalJSON(b []byte) error {
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
