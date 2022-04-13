package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

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

func (c UrlDomainMinSslVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UrlDomainMinSslVersion) UnmarshalJSON(b []byte) error {
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
