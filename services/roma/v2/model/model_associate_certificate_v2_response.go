package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AssociateCertificateV2Response Response Object
type AssociateCertificateV2Response struct {

	// 自定义域名
	UrlDomain string `json:"url_domain"`

	// 自定义域名的编号
	Id string `json:"id"`

	// CNAME解析状态 - 1: 未解析 - 2: 解析中 - 3: 解析成功 - 4: 解析失败
	Status AssociateCertificateV2ResponseStatus `json:"status"`

	// 支持的最小SSL版本
	MinSslVersion string `json:"min_ssl_version"`

	// 是否开启http到https的重定向，false为关闭，true为开启，默认为false
	IsHttpRedirectToHttps *bool `json:"is_http_redirect_to_https,omitempty"`

	// 是否开启客户端证书校验。只有绑定证书时，该参数才生效。当绑定证书存在trusted_root_ca时，默认开启；当绑定证书不存在trusted_root_ca时，默认关闭。
	VerifiedClientCertificateEnabled *bool `json:"verified_client_certificate_enabled,omitempty"`

	// 证书的名称
	SslName *string `json:"ssl_name,omitempty"`

	// 证书的编号
	SslId *string `json:"ssl_id,omitempty"`

	// SSL证书列表  [暂不支持](tag:hws;hws_hk;fcs;g42;Site)
	SslInfos       *[]SslInfo `json:"ssl_infos,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o AssociateCertificateV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateCertificateV2Response struct{}"
	}

	return strings.Join([]string{"AssociateCertificateV2Response", string(data)}, " ")
}

type AssociateCertificateV2ResponseStatus struct {
	value int32
}

type AssociateCertificateV2ResponseStatusEnum struct {
	E_1 AssociateCertificateV2ResponseStatus
	E_2 AssociateCertificateV2ResponseStatus
	E_3 AssociateCertificateV2ResponseStatus
	E_4 AssociateCertificateV2ResponseStatus
}

func GetAssociateCertificateV2ResponseStatusEnum() AssociateCertificateV2ResponseStatusEnum {
	return AssociateCertificateV2ResponseStatusEnum{
		E_1: AssociateCertificateV2ResponseStatus{
			value: 1,
		}, E_2: AssociateCertificateV2ResponseStatus{
			value: 2,
		}, E_3: AssociateCertificateV2ResponseStatus{
			value: 3,
		}, E_4: AssociateCertificateV2ResponseStatus{
			value: 4,
		},
	}
}

func (c AssociateCertificateV2ResponseStatus) Value() int32 {
	return c.value
}

func (c AssociateCertificateV2ResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AssociateCertificateV2ResponseStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
